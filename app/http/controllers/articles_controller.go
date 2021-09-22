package controllers

import (
	"database/sql"
	"fmt"
	"goblong/pkg/logger"
	"goblong/pkg/route"
	"goblong/pkg/types"
	"html/template"
	"net/http"
)

type ArticlesController struct {
}

func (*ArticlesController) Show(w http.ResponseWriter, r *http.Request) {

	// 1. get url parameters
	id := route.GetRouteVariable("id", r)

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
