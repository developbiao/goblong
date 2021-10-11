package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/thedevsaddam/govalidator"
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
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	// 2. Validation rules
	rules := govalidator.MapData{
		"name":             []string{"required", "alpha_num", "between:3,20"},
		"email":            []string{"required", "min:4", "max:30", "email"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}

	// 2.2 Custom valida message
	messages := govalidator.MapData{
		"name": []string{
			"required:用户名是必填项",
			"alpha_num:格式错误，只允许数字和英文",
			"between:用户名是3~20之间",
		},
		"email": []string{
			"required:Email是必填项",
			"min:Email长度需大于6",
			"max:Email长度需小于30",
		},
		"password": []string{
			"required:密码是必填项",
			"min:长度需要大于或等于6",
		},
		"password_confirm": []string{
			"required:确认密码是必填项",
		},
	}

	// 3. config option
	opts := govalidator.Options{
		Data:          &_user,
		Rules:         rules,
		TagIdentifier: "valid", // struct identifier
		Messages:      messages,
	}

	// 4. start  validation
	errs := govalidator.New(opts).ValidateStruct()
	if len(errs) > 0 {
		// error happen detected
		data, _ := json.MarshalIndent(errs, "", " ")
		fmt.Fprint(w, string(data))
	} else {
		// 2. create user and redirect to home page
		//_user := user.User{
		//	Name:     name,
		//	Email:    email,
		//	Password: password,
		//}
		//
		//_user.Create()
		//
		//if _user.ID > 0 {
		//	fmt.Fprint(w, "Insert user success, ID is "+_user.GetStringID())
		//} else {
		//	w.WriteHeader(http.StatusInternalServerError)
		//	fmt.Print(w, "Create user failed, Please contact administrator")
		//}

		fmt.Fprint(w, "Validation success!")

	}

	// 3. invalid form re display register form page
}
