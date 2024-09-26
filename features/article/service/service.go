package service

import (
	"qhealth/domain"
	"qhealth/features/article"
)

type service struct {
	repo article.Repository
}

func NewArticleService(repo article.Repository) article.Service {
	return &service{repo: repo}
}

func (s *service) CreateArticle(articleReq domain.ArticleReq, userId string) error {
	article := domain.ReqToArticle(articleReq, userId)

	err := s.repo.CreateArticle(article)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetUserByEmail(email string) (domain.User, error) {
    return s.repo.GetUserByEmail(email)
}


func (s *service) GetAllArticle(title string, userId string) ([]domain.ArticleResp, error) {
	article, err := s.repo.GetAllArticle(title)

	if err != nil {
		return nil, err
	}

	result := domain.ListArticleToResp(article)
	return result, nil
}

func (s *service) GetLatestArticle() ([]domain.ArticleResp, error) {
	article, err := s.repo.GetLatestArticle()

	if err != nil {
		return nil, err
	}

	result := domain.ListArticleToResp(article)
	return result, nil
}

func (s *service) GetArticleById(id string) (*domain.Articles, error) {
	return s.repo.GetArticleById(id)
}

func (s *service) UpdateArticle(id string, article domain.ArticleReq) error {
	data := domain.ReqToArticle(article, "")

	err := s.repo.UpdateArticle(id, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) DeleteArticle(id string) error {
	return s.repo.DeleteArticle(id)
}