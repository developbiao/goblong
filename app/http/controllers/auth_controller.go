package controllers

import (
	"encoding/json"
	"fmt"
	"goblong/app/models/user"
	"goblong/app/requests"
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
		data, _ := json.MarshalIndent(errs, "", " ")
		fmt.Fprint(w, string(data))
	} else {
		//  create user and redirect to home page
		_user.Create()

		if _user.ID > 0 {
			fmt.Fprint(w, "Insert user success, ID is "+_user.GetStringID())
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Print(w, "Create user failed, Please contact administrator")
		}

		fmt.Fprint(w, "Validation success!")

	}
	//  invalid form re display register form page
}
