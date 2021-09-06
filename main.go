package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

var router = mux.NewRouter()

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1 style=\"color:pink\">Hello, welcome my goblog</h1>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>This blog just for record leanring golang, if you have any question please contact "+
		"<a href=\"mailto:developbiao@gmail.com\">developbiao@gmail.com</a></h1>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Reqeust not found page :(</h1>"+
		"<p>If you have any doubts, please contact us. </p>")

}

func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprint(w, "Article ID: "+id)
}

// Store article information
func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprint(w, "Please provide correct data")
		return
	}

	title := r.PostForm.Get("title")

	fmt.Fprintf(w, "POST PostFrom: %v <br/>", r.PostForm)
	fmt.Fprintf(w, "POST From: %v <br/>", r.Form)
	fmt.Fprintf(w, "title value: %v", title)

	fmt.Fprintf(w, "r From title value: %v <br/>", r.FormValue("title"))
	fmt.Fprintf(w, "r PostFrom title value: %v <br/>", r.PostFormValue("title"))
	fmt.Fprintf(w, "r Form agent value: %v<br/>", r.FormValue("agent"))
	fmt.Fprintf(w, "r PostForm agent value: %v<br/>", r.FormValue("agent"))
}

func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Visit article")
}

// force add html header
func forceHTMLMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 1. Set header
		w.Header().Set("Content-Type", "text/html; charset=utf-8")

		// 2. Continue request
		next.ServeHTTP(w, r)
	})
}

// Remove trailing slash
func removeTrailingSlash(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}
		// continue serve
		next.ServeHTTP(w, r)
	})
}

// Create article handler
func articlesCreateHandler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="en">
<head>
<title>Create a new blog </title>
</head>
<body>
	<form action="%s?agent=proxy" method="post">
		<p><input type="text" name="title"></p>
		<p><textarea name="body" cols="30" rows="10"></textarea></p>
		<p><button type="submit">提交</button></p>
	</form>
</body>
</html>
`
	storeURL, _ := router.Get("articles.store").URL()
	fmt.Fprintf(w, html, storeURL)

}

func main() {
	router.HandleFunc("/", homeHandler).Methods("GET").Name("home")
	router.HandleFunc("/about", aboutHandler).Methods("GET").Name("about")

	// Articles router
	router.HandleFunc("/articles/{id:[0-9]+}",
		articlesShowHandler).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", articlesCreateHandler).
		Methods("GET").
		Name("articles.create")

	// Custom 404 page
	router.NotFoundHandler = http.HandlerFunc(notFoundHandler)

	// Get router name URL example
	homeURL, _ := router.Get("home").URL()
	fmt.Println("homeURL: ", homeURL)
	articleURL, _ := router.Get("articles.show").URL("id", "23")
	fmt.Println("articleURL: ", articleURL)

	router.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprint(w, "From [GET] request ")
		case "POST":
			fmt.Fprint(w, "From [POST] request")
		}
	})

	// Middleware force content is HTML
	router.Use(forceHTMLMiddleware)

	http.ListenAndServe(":3000", removeTrailingSlash(router))
}
