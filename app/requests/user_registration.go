package requests

import (
	"goblong/app/models/user"

	"github.com/thedevsaddam/govalidator"
)

func ValidateRegistrationForm(data user.User) map[string][]string {

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
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid", // struct identifier
		Messages:      messages,
	}

	// 4. start  validation
	errs := govalidator.New(opts).ValidateStruct()

	// 5. validation password_confirm
	if data.Password != data.PasswordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "两次输入密码不匹配")
	}

	return errs
}