package article

import (
	"goblong/pkg/model"
	"goblong/pkg/types"
)

// Article model
type Article struct {
	ID    int
	Title string
	Body  string
}

// Get article by id
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)
	if err := model.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}
