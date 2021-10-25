package policies

import (
	"goblong/app/models/article"
	"goblong/pkg/auth"
)

//  Check can modify topic
func CanModifyArtile(_article article.Article) bool {
	return auth.User().ID == _article.UserID
}
