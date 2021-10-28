package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblong/app/models/category"
)

func ValidateCategoryForm(data category.Category) map[string][]string {
	// Define rules
	rules := govalidator.MapData{
		"name": []string{"required", "min_cn:2", "max_cn:8", "not_exists:categories,name"},
	}

	// Define error message
	messages := govalidator.MapData{
		"name": []string{
			"required:category name is required",
			"min_cn:category length at least 2",
			"max_cn:category length must less than 8",
		},
	}

	// Option init
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}
	return govalidator.New(opts).ValidateStruct()

}
