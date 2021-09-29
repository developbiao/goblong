package controllers

import (
	"database/sql"
	"fmt"
	"goblong/pkg/logger"
	"goblong/pkg/model/article"
	"goblong/pkg/route"
	"goblong/pkg/types"
	"gorm.io/gorm"
	"html/template"
	"net/http"
	"path/filepath"
	"strconv"
	"unicode/utf8"
)

type ArticlesController struct {
}

// Articles from data
type ArticlesFormData struct {
	Title, Body string
	URL         string
	Errors      map[string]string
}

// Show article
func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {

	// 1. get url parameters
	id := route.GetRouteVariable("id", r)

	// Get record by articleRecord id
	articleRecord, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// Not found record
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 articleRecord not found")
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
		tmpl.Execute(w, articleRecord)
		if err != nil {
			logger.LogError(err)
		}

	}

}

// Articles
func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {

	// Get articles
	articles, err := article.GetAll()

	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 Internal Server error")

	} else {
		// Load templates

		// Set template related path
		viewDir := "resources/views"

		// All layout slices
		files, err := filepath.Glob(viewDir + "/layouts/*.gohtml")
		logger.LogError(err)

		// Add our target file
		newFiles := append(files, viewDir+"/articles/index.gohtml")

		// Parsing template file
		tmpl, err := template.ParseFiles(newFiles...)
		logger.LogError(err)

		// Render template files
		tmpl.ExecuteTemplate(w, "app", articles)

	}

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

// Create article page
func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	storeURL := route.Name2URL("articles.store")
	data := ArticlesFormData{
		Title:  "",
		Body:   "",
		URL:    storeURL,
		Errors: nil,
	}

	tmpl, err := template.ParseFiles("resources/views/articles/create.gohtml")
	if err != nil {
		panic(err)
	}
	tmpl.Execute(w, data)
}

// Store article
func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {

	title := r.PostFormValue("title")
	body := r.PostFormValue("body")

	errors := validateArticleFormData(title, body)

	if len(errors) == 0 {
		_article := article.Article{
			Title: title,
			Body:  body,
		}
		_article.Create()
		if _article.ID > 0 {
			fmt.Fprint(w, "Insert ID: "+strconv.FormatInt(int64(_article.ID), 10))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Create article failed please contact administrator")
		}

	} else {

		storeURL := route.Name2URL("article.store")
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

// Edit article
func (*ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {

	// get url parameters
	id := route.GetRouteVariable("id", r)

	// Get record by article id
	articleRecord, err := article.Get(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		updateURL := route.Name2URL("articles.update", "id", id)
		data := ArticlesFormData{
			Title:  articleRecord.Title,
			Body:   articleRecord.Body,
			URL:    updateURL,
			Errors: nil,
		}
		tmpl, err := template.ParseFiles("resources/views/articles/edit.gohtml")
		logger.LogError(err)

		tmpl.Execute(w, data)
	}
}

// Update article
func (*ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	// Get url parameter
	id := route.GetRouteVariable("id", r)

	// Get article
	_article, err := article.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
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
			_article.Title = title
			_article.Body = body
			rowsAffected, err := _article.Update()
			if err != nil {
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 Internal Server Error")
			}

			if rowsAffected > 0 {
				// Update success
				showURL := route.Name2URL("articles.show", "id", id)
				if err != nil {
					fmt.Fprint(w, "Sorry, you don't have permission~", err)
				}
				http.Redirect(w, r, showURL, http.StatusFound)
			} else {
				fmt.Fprint(w, "You not change anything~")
			}

		} else {
			updateURL := route.Name2URL("articles.update", "id", id)
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

// Delete action
func (*ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {
	// Get id
	id := route.GetRouteVariable("id", r)

	// Get article by id
	_article, err := article.Get(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 not found article")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal server error")
		}

	} else {
		rowsAffected, err := _article.Delete()
		if err != nil {
			// Should be sql error
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Internal server error")
		}

		if rowsAffected > 0 {
			indexURL := route.Name2URL("articles.index")
			http.Redirect(w, r, indexURL, http.StatusFound)

		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 not found article")
		}

	}

}
