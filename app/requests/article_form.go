package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblong/app/models/article"
)

// Validate article form
func ValidateArticleForm(data article.Article) map[string][]string {
	// 1. defined article role
	rules := govalidator.MapData{
		"title": []string{"required", "min:3", "max:40"},
		"body":  []string{"required", "min:10"},
	}

	// 2. defined error message
	messages := govalidator.MapData{
		"title": []string{
			"required:article title is required",
			"min:article title length must be granter than 3",
			"max:article title length must be less than 40",
		},
		"body": []string{
			"required:article body is required",
			"min:article body must be granter than 10",
		},
	}

	// 3. Config validator options
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	// start validation
	return govalidator.New(opts).ValidateStruct()
}
