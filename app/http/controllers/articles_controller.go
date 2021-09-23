package controllers

import (
	"fmt"
	"goblong/pkg/logger"
	"goblong/pkg/model/article"
	"goblong/pkg/route"
	"goblong/pkg/types"
	"gorm.io/gorm"
	"html/template"
	"net/http"
)

type ArticlesController struct {
}

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
		// Load template
		tmpl, err := template.ParseFiles("resources/views/articles/index.gohtml")
		logger.LogError(err)

		// Render template
		tmpl.Execute(w, articles)

	}

}
