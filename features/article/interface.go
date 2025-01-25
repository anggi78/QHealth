package article

import (
	"qhealth/domain"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		CreateArticle(article domain.Articles) error
		GetUserByEmail(email string) (domain.User, error)
		GetAllArticle(title string, page, pageSize int) ([]domain.Articles, int, error)
		GetLatestArticle() ([]domain.Articles, error)
		GetArticleById(id string) (*domain.Articles, error)
		UpdateArticle(id string, article domain.Articles) error
		DeleteArticle(id string) error
	}

	Service interface {
		CreateArticle(articleReq domain.ArticleReq) error
		GetUserByEmail(email string) (domain.User, error)
		GetAllArticle(title, userId string, page, pageSize int) ([]domain.ArticleResp, int, error)
		GetLatestArticle() ([]domain.ArticleResp, error)
		GetArticleById(id string) (*domain.Articles, error)
		UpdateArticle(id string, article domain.ArticleReq) error
		DeleteArticle(id string) error
	}

	Handler interface {
		CreateArticle(e echo.Context) error
		GetAllArticle(e echo.Context) error
		GetLatestArticle(e echo.Context) error
		GetArticleById(e echo.Context) error
		UpdateArticle(e echo.Context) error
		DeleteArticle(e echo.Context) error
	}
)