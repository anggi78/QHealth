package repository

import (
	"qhealth/domain"
	articleview "qhealth/features/article_view"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewViewRepository(db *gorm.DB) articleview.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetAllView() ([]domain.ArticleView, error) {
    var views []domain.ArticleView

    err := r.db.Preload("User").Preload("Article").Find(&views).Error
    if err != nil {
        return nil, err
    }

    return views, nil
}

func (r *repository) GetArticleById(articleId string) (domain.Articles, error) {
	var article domain.Articles

	err := r.db.First(&article, "id = ?", articleId).Error
	return article, err
}

func (r *repository) CreateArticleView(articleView domain.ArticleView) error {
	return r.db.Create(&articleView).Error
}
