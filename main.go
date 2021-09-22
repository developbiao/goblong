package main

import (
	"database/sql"
	"fmt"
	"goblong/bootstrap"
	"goblong/pkg/database"
	"goblong/pkg/logger"
	"goblong/pkg/route"
	"goblong/pkg/types"
	"html/template"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"unicode/utf8"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

//var router = mux.NewRouter()
var router *mux.Router
var db *sql.DB

// ArticlesFormData struct
type ArticlesFormData struct {
	Title, Body string
	URL         *url.URL
	Errors      map[string]string
}

// Article record
type Article struct {
	Title, Body string
	ID          int64
}

// Show article by id
func articlesShowHandler(w http.ResponseWriter, r *http.Request) {
	// 1. get url parameters
	id := getRouteVariable("id", r)

	// Get record by article id
	article, err := getArticleByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			// Not found record
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 article not found")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Server Internal error")
		}
	} else {
		tmpl, err := template.New("show.gohtml").
			Funcs(template.FuncMap{
				"RouteName2URL": route.Name2URL,
				"Int64ToString": types.Int64ToString,
			}).
			ParseFiles("resources/views/articles/show.gohtml")
		if err != nil {
			logger.LogError(err)
		}
		// Read success
		tmpl.Execute(w, article)
		if err != nil {
			logger.LogError(err)
		}

	}
}

// Store article information
func articlesStoreHandler(w http.ResponseWriter, r *http.Request) {
	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	errors := validateArticleFormData(title, body)

	if len(errors) == 0 {
		lastInsertId, err := saveArticleToDB(title, body)
		if lastInsertId > 0 {
			fmt.Fprint(w, "Insert ID: "+strconv.FormatInt(lastInsertId, 10))
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Server Internal error")
		}

	} else {

		storeURL, _ := router.Get("articles.store").URL()
		data := ArticlesFormData{
			Title:  title,
			Body:   body,
			URL:    storeURL,
			Errors: errors,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
		if err != nil {
			panic(err)
		}

		tmpl.Execute(w, data)
	}

}

func saveArticleToDB(title string, body string) (int64, error) {
	// Init variables
	var (
		id   int64
		err  error
		rs   sql.Result
		stmt *sql.Stmt
	)

	// Get a prepare
	stmt, err = db.Prepare("INSERT INTO `articles` (`title`, `body`) VALUES(?, ?)")
	// Check error
	if err != nil {
		return 0, err
	}

	//defer db.Close()

	rs, err = stmt.Exec(title, body)
	if err != nil {
		return 0, err
	}

	id, err = rs.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Articles index handler
func articlesIndexHandler(w http.ResponseWriter, r *http.Request) {
	// Get query result
	rows, err := db.Query("SELECT * FROM `articles`")
	logger.LogError(err)
	defer rows.Close()

	var articles []Article
	for rows.Next() {
		var article Article
		err := rows.Scan(&article.ID, &article.Title, &article.Body)
		logger.LogError(err)
		// Append article to articles slice
		articles = append(articles, article)
	}

	// Check iterator error
	err = rows.Err()
	logger.LogError(err)

	// Load template
	tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
	logger.LogError(err)

	// Render template
	tmpl.Execute(w, articles)
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

// Article edit handler
func articlesEditHandler(w http.ResponseWriter, r *http.Request) {
	// get url parameters
	id := getRouteVariable("id", r)

	// Get record by article id
	article, err := getArticleByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		updateURL, _ := router.Get("articles.update").URL("id", id)
		data := ArticlesFormData{
			Title:  article.Title,
			Body:   article.Body,
			URL:    updateURL,
			Errors: nil,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
		logger.LogError(err)

		tmpl.Execute(w, data)
	}

}

// Article update handler
func articlesUpdateHandler(w http.ResponseWriter, r *http.Request) {
	// Get url parameter
	id := getRouteVariable("id", r)

	// Get article
	_, err := getArticleByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			// Not found data
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 article not found")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal server error")
		}
	} else {
		title := r.PostFormValue("title")
		body := r.PostFormValue("body")

		errors := validateArticleFormData(title, body)
		if len(errors) == 0 {
			query := "UPDATE `articles` SET `title` = ?, body = ? WHERE `id` = ?"
			rs, err := db.Exec(query, title, body, id)
			if err != nil {
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 Internal Server Error")
			}

			if n, _ := rs.RowsAffected(); n > 0 {
				// Update success
				showURL, _ := router.Get("articles.show").URL("id", id)
				if err != nil {
					fmt.Fprint(w, "Sorry, you don't have permission~", err)
				}
				http.Redirect(w, r, showURL.String(), http.StatusFound)
			} else {
				fmt.Fprint(w, "You not change anything~")
			}

		} else {
			updateURL, _ := router.Get("articles.update").URL("id", id)
			data := ArticlesFormData{
				Title:  title,
				Body:   body,
				URL:    updateURL,
				Errors: errors,
			}

			tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
			logger.LogError(err)
			tmpl.Execute(w, data)
		}

	}
}

// Get article by id
func getArticleByID(id string) (Article, error) {
	article := Article{}
	query := "SELECT * FROM `articles` WHERE `id` = ?"
	err := db.QueryRow(query, id).Scan(&article.ID, &article.Title, &article.Body)
	return article, err
}

// Validation article form data
func validateArticleFormData(title string, body string) map[string]string {

	errors := make(map[string]string)

	// Validation title
	if title == "" {
		errors["title"] = "Title can'not is empty"
	} else if utf8.RuneCountInString(title) < 3 || utf8.RuneCountInString(title) > 40 {
		errors["title"] = "Title length must between 3 ~ 40 characters"
	}

	// Validation  body
	if body == "" {
		errors["body"] = "Body can'not is empty"
	} else if utf8.RuneCountInString(body) < 10 {
		errors["body"] = "Body length must granter than 10 characters"
	}
	return errors
}

// Generate link
func (a Article) Link() string {
	showURL, err := router.Get("articles.show").URL("id", strconv.FormatInt(a.ID, 10))
	if err != nil {
		logger.LogError(err)
		return ""
	}
	return showURL.String()
}

// Delete article
func (a Article) Delete() (rowsAffected int64, err error) {
	rs, err := db.Exec("DELETE FROM `articles` WHERE `id` = ?", a.ID)
	if err != nil {
		return 0, err
	}

	if n, _ := rs.RowsAffected(); n > 0 {
		return n, nil
	}
	return 0, nil
}

// Article delete
func articlesDeleteHandler(w http.ResponseWriter, r *http.Request) {
	// Get id
	id := getRouteVariable("id", r)

	// Get article by id
	article, err := getArticleByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 not found article")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal server error")
		}

	} else {
		rowsAffected, err := article.Delete()
		if err != nil {
			// Should be sql error
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal server error")
		}

		if rowsAffected > 0 {
			indexURL, _ := router.Get("articles.index").URL()
			http.Redirect(w, r, indexURL.String(), http.StatusFound)

		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 not found article")
		}

	}

}

func getRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}

func main() {

	// initialize database
	database.Initialize()
	db = database.DB

	// initialize router
	router = bootstrap.SetupRoute()

	// Articles index
	router.HandleFunc("/articles", articlesIndexHandler).Methods("GET").Name("articles.index")
	// Create article page
	router.HandleFunc("/articles/create", articlesCreateHandler).
		Methods("GET").
		Name("articles.create")

	// Save article
	router.HandleFunc("/articles", articlesStoreHandler).Methods("POST").Name("articles.store")

	// edit
	router.HandleFunc("/articles/{id:[0-9]+}/edit", articlesEditHandler).
		Methods("GET").
		Name("articles.edit")
	// update
	router.HandleFunc("/articles/{id:[0-9]+}", articlesUpdateHandler).
		Methods("POST").
		Name("articles.update")

	// Delete
	router.HandleFunc("/articles/{id:[0-9]+}/delete", articlesDeleteHandler).
		Methods("POST").
		Name("articles.delete")

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
