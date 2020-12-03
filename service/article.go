package service

import (
	"mermaid/model"
)

type article struct{}

var DefaultArticle = &article{}

func (a *article) GetAllArticles() ([]*model.Articles, error) {
	db := model.DB
	data := make([]*model.Articles, 0)
	err := db.Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
