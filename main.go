package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"net/url"
	"strings"
	"unicode/utf8"
)

var router = mux.NewRouter()

// ArticlesFormData struct
type ArticlesFormData struct {
	Title, Body string
	URL         *url.URL
	Errors      map[string]string
}

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
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	errors := make(map[string]string)

	// Validation title
	if title == "" {
		errors["title"] = "Title can'not empty"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "Title length between 3~40"
	}

	// Validation body
	if body == "" {
		errors["body"] = "Body can'not empty"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "Body length must great than 10 character"
	}

	if len(errors) == 0 {
		fmt.Fprint(w, "Valid!")
		fmt.Fprintf(w, "title value: %v \n", title)
		fmt.Fprintf(w, "title length: %v \n", utf8.RuneCountInString(title))
		fmt.Fprintf(w, "body value: %v \n", body)
		fmt.Fprintf(w, "body length: %v \n", utf8.RuneCountInString(body))
	} else {

		html := `
<!DOCTYPE html>
<html lang="en">
<head>
    <title>创建文章 —— 我的技术博客</title>
    <style type="text/css">.error {color: red;}</style>
</head>
<body>
    <form action="{{ .URL }}" method="post">
        <p><input type="text" name="title" value="{{ .Title }}"></p>
        {{ with .Errors.title }}
        <p class="error">{{ . }}</p>
        {{ end }}
        <p><textarea name="body" cols="30" rows="10">{{ .Body }}</textarea></p>
        {{ with .Errors.body }}
        <p class="error">{{ . }}</p>
        {{ end }}
        <p><button type="submit">提交</button></p>
    </form>
</body>
</html>
`
		storeURL, _ := router.Get("articles.store").URL()
		data := ArticlesFormData{
			Title:  title,
			Body:   body,
			URL:    storeURL,
			Errors: errors,
		}
		tmpl, err := template.New("create-form").Parse(html)
		if err != nil {
			panic(err)
		}

		tmpl.Execute(w, data)
	}

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
