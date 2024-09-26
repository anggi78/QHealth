package repository

import (
	"qhealth/domain"
	"qhealth/features/article"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) article.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) CreateArticle(article domain.Articles) error {
	err := r.db.Create(&article).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetUserByEmail(email string) (domain.User, error) {
    var user domain.User
    if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
        return user, err
    }
    return user, nil
}

func (r *repository) GetAllArticle(title string) ([]domain.Articles, error) {
	var article []domain.Articles
	query := r.db.Order("created_at")

	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	err := query.Preload("User").Find(&article).Error
	if err != nil {
		return nil, err
	}

	return article, nil
}

func (r *repository) GetLatestArticle() ([]domain.Articles, error) {
	var latestArticle []domain.Articles
	query := r.db.Order("created_at desc").Limit(10)

	err := query.Find(&latestArticle).Error
	if err != nil {
		return nil, err
	}

	return latestArticle, nil
}

func (r *repository) GetArticleById(id string) (*domain.Articles, error) {
	var article domain.Articles

	err := r.db.First(&article, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (r *repository) UpdateArticle(id string, article domain.Articles) error {
	_, err := r.GetArticleById(id)
	if err != nil {
		return err
	}

	err = r.db.Where("id = ?", id).Updates(&article).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *repository) DeleteArticle(id string) error {
	err := r.db.Delete(&domain.Articles{}, "id = ?", id).Error
	if err != nil {
		return err
	}

	return nil
}