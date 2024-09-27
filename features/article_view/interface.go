package articleview

import (
	"qhealth/domain"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		GetAllView() ([]domain.ArticleView, error)
		GetArticleById(articleId string) (domain.Articles, error)
		CreateArticleView(articleView domain.ArticleView) error
	}

	Service interface {
		GetArticleTrackView(userId, articleId string) (domain.Articles, error)
		GetAllView() ([]domain.ArticleViewResp, error)
	}

	Handler interface {
		GetAllView(e echo.Context) error
		GetArticleTrackView(e echo.Context) error
	}
)