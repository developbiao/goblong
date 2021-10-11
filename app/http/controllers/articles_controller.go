package controllers

import (
	"fmt"
	"goblong/app/models/article"
	"goblong/pkg/logger"
	"goblong/pkg/route"
	"goblong/pkg/view"
	"gorm.io/gorm"
	"net/http"
	"unicode/utf8"
)

type ArticlesController struct {
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
		// Render article
		view.Render(w, articleRecord, "articles.show")

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
		// Render articles
		view.Render(w, articles, "articles.index")
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
	view.Render(w, view.D{}, "articles.create", "articles._form_field")
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
			fmt.Fprint(w, "Insert ID: "+_article.GetStringID())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Create article failed please contact administrator")
		}

	} else {
		view.Render(w, view.D{
			"Title":  title,
			"Body":   body,
			"Errors": errors,
		}, "articles.create", "articles._form_field")

	}

}

// Edit article
func (*ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {

	// get url parameters
	id := route.GetRouteVariable("id", r)

	// Get record by article id
	_articleRecord, err := article.Get(id)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		view.Render(w, view.D{
			"Article": _articleRecord,
			"Errors":  view.D{},
		}, "articles.edit", "articles._form_field")
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
			view.Render(w, view.D{
				"Article": _article,
				"Errors":  errors,
			}, "articles.edit", "articles._form_field")
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
