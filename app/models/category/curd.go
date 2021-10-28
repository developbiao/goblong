package category

import (
	"goblong/pkg/logger"
	"goblong/pkg/model"
	"goblong/pkg/types"
)

// Create category record
func (category *Category) Create() (err error) {
	if err = model.DB.Create(&category).Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// Get all categories
func All() ([]Category, error) {
	var categories []Category
	if err := model.DB.Find(&categories).Error; err != nil {
		logger.LogError(err)
		return categories, err
	}
	return categories, nil
}

// Get category record by id
func Get(idstr string) (Category, error) {
	var category Category
	id := types.StringToInt(idstr)
	if err := model.DB.First(&category, id).Error; err != nil {
		return category, err
	}
	return category, nil
}
