package article

import (
	"goblong/app/models"
	"goblong/pkg/route"
)

// Article model
type Article struct {
	models.BaseModel

	Title string
	Body  string
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}
