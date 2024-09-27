package router

import (
	"qhealth/features/article_view/handler"
	"qhealth/features/article_view/repository"
	"qhealth/features/article_view/service"
	"qhealth/helpers/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ViewRoute(v *echo.Group, db *gorm.DB) {
	repo := repository.NewViewRepository(db)
	serv := service.NewViewService(repo)
	handler := handler.NewViewHandler(serv)

	v.GET("", handler.GetAllView)

	article := v.Group("/article", middleware.JwtMiddleware())
	article.GET("/:id", handler.GetArticleTrackView)
}