package controllers

import (
	"fmt"
	"goblong/app/models/user"
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
	// 0. get initialization form variables
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	// 1. validation form

	// 2. create user and redirect to home page
	_user := user.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	_user.Create()

	if _user.ID > 0 {
		fmt.Fprint(w, "Insert user success, ID is "+_user.GetStringID())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Print(w, "Create user failed, Please contact administrator")
	}

	// 3. invalid form re display register form page
}
