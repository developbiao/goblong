package article

import (
	"goblong/app/models"
	"goblong/app/models/user"
	"goblong/pkg/route"
)

// Article model
type Article struct {
	models.BaseModel

	Title  string `gorm:"type:varchar(255);not null;" valid:"title"`
	Body   string `gorm:"type:longtext;not null;" valid:"body"`
	UserID uint64 `gorm:"not null;index"`
	User   user.User
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}

func (a Article) CreatedAtDate() string {
	return a.CreatedAt.Format("2006-01-02")
}
