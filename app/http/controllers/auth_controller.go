package controllers

import (
	"goblong/pkg/view"
	"net/http"
)

// AuthorController process static page
type AuthController struct {
}

// Register page
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// Do Register logic
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	//
}
