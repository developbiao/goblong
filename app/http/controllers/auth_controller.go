package controllers

import (
	"fmt"
	"goblong/app/models/user"
	"goblong/app/requests"
	"goblong/pkg/auth"
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
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	// 4. start  validation
	errs := requests.ValidateRegistrationForm(_user)
	if len(errs) > 0 {
		// error happen detected
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")
	} else {
		//  create user and redirect to home page
		_user.Create()

		if _user.ID > 0 {
			fmt.Fprint(w, "Insert user success, ID is "+_user.GetStringID())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Create user failed, Please contact administrator")
		}

	}
	//  invalid form re display register form page
}

// Login page
func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.login")
}

// Process post login
func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {
	// initialization form
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	// Attempt login
	if err := auth.Attempt(email, password); err == nil {
		// Login success
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// Login failed display error
		view.RenderSimple(w, view.D{
			"Error":    err.Error(),
			"Email":    email,
			"Password": password,
		}, "auth.login")
	}

}

// Logout
func (*AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	if auth.Check() {
		auth.Logout()
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
