package controllers

import (
	"fmt"
	"goblong/app/models/article"
	"goblong/app/models/category"
	"goblong/app/requests"
	"goblong/pkg/route"
	"goblong/pkg/view"
	"net/http"
)

type CategoriesController struct {
	BaseController
}

// Create article category page
func (*CategoriesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "categories.create")
}

// Store article category
func (*CategoriesController) Store(w http.ResponseWriter, r *http.Request) {
	// init category
	_category := category.Category{
		Name: r.PostFormValue("name"),
	}

	// Validation
	errors := requests.ValidateCategoryForm(_category)
	if len(errors) != 0 {
		view.Render(w, view.D{
			"Category": _category,
			"Errors":   errors,
		}, "categories.create")

	} else {
		// Create article
		_category.Create()
		if _category.ID > 0 {
			fmt.Fprint(w, "Create success!")
			//indexURL := route.Name2URL("categories.show", "id", _category.GetStringID())
			//http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			fmt.Println("Create category failed!")
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (cc *CategoriesController) Show(w http.ResponseWriter, r *http.Request) {
	// 1. Get url parameter
	id := route.GetRouteVariable("id", r)

	// 2. Get category from id
	_category, err := category.Get(id)
	if err != nil {
		cc.ResponseFromSQLError(w, err)
		return
	}

	// 3. Get result
	articles, pagerData, err := article.GetByCategoryID(_category.GetStringID(), r, 5)
	if err != nil {
		cc.ResponseFromSQLError(w, err)
	} else {
		// Render template
		view.Render(w, view.D{
			"Articles":  articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}

}
