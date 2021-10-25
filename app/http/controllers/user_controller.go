package controllers

import (
	"fmt"
	"goblong/app/models/article"
	"goblong/app/models/user"
	"goblong/pkg/logger"
	"goblong/pkg/route"
	"goblong/pkg/view"
	"net/http"

	"gorm.io/gorm"
)

// User controller
type UserController struct {
}

// Show user prifile page
func (*UserController) Show(w http.ResponseWriter, r *http.Request) {
	// Get url parameter
	id := route.GetRouteVariable("id", r)

	// Get articles by user id
	_user, err := user.Get(id)

	// Check error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 user not found")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal server error")
		}
	} else {
		articles, err := article.GetByUserID(_user.GetStringID())
		if err != nil {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 Internal server error")
		} else {
			view.Render(w, view.D{
				"Articles": articles,
			}, "articles.index", "articles._article_meta")
		}
	}

}
