package article

import (
	"goblong/pkg/route"
	"strconv"
)

// Article model
type Article struct {
	ID    int
	Title string
	Body  string
}

func (a Article) Link() string {
	return route.Name2URL("articles.show", "id", strconv.FormatInt(int64(a.ID), 10))
}
