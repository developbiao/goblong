package article

import (
	"goblong/pkg/logger"
	"goblong/pkg/model"
	"goblong/pkg/pagination"
	"goblong/pkg/route"
	"goblong/pkg/types"
	"net/http"
)

// Get article by id
func Get(idstr string) (Article, error) {
	var article Article
	id := types.StringToInt(idstr)
	if err := model.DB.Preload("User").First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

// Get All articles
func GetAll(r *http.Request, perPage int) ([]Article, pagination.ViewData, error) {
	// 1. init pager instance
	db := model.DB.Model(Article{}).Order("created_at DESC")
	_pager := pagination.New(r, db, route.Name2URL("articles.index"), perPage)

	// 2. Get view data
	viewData := _pager.Paging()

	var articles []Article
	_pager.Results(&articles)
	return articles, viewData, nil
}

// Get articles by user id
func GetByUserID(uid string) ([]Article, error) {
	var articles []Article
	if err := model.DB.Where("`user_id` = ?", uid).Preload("User").Order("updated_at DESC").Find(&articles).Error; err != nil {
		return articles, err
	}
	return articles, nil
}

// Create article
func (article *Article) Create() (err error) {
	result := model.DB.Create(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return err
	}
	return nil
}

// Update article
func (article *Article) Update() (rowsAffected int64, err error) {
	result := model.DB.Save(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}

// Delete article
func (article *Article) Delete() (rowsAffected int64, err error) {
	result := model.DB.Delete(&article)
	if err = result.Error; err != nil {
		logger.LogError(err)
		return 0, err
	}
	return result.RowsAffected, nil
}
