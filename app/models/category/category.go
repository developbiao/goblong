package category

import (
	"goblong/app/models"
	"goblong/pkg/route"
)

type Category struct {
	models.BaseModel

	Name string `gorm:"type:varchar(255);not null;" valid:"name"`
}

// Link generator
func (category Category) Link() string {
	return route.Name2URL("categories.show", "id", category.GetStringID())
}
