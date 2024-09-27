package router

import (
	"qhealth/features/article/handler"
	"qhealth/features/article/repository"
	"qhealth/features/article/service"
	"qhealth/helpers/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ArticleRoute(a *echo.Group, db *gorm.DB) {
	repo := repository.NewArticleRepository(db)
	serv := service.NewArticleService(repo)
	handler := handler.NewArticleHandler(serv)

	a.GET("/latest", handler.GetLatestArticle)

	admin := a.Group("/admin", middleware.JwtMiddleware())
	admin.GET("", handler.GetAllArticle)
	admin.GET("/:id", handler.GetArticleById)
	admin.POST("", handler.CreateArticle)
	admin.PUT("/:id", handler.UpdateArticle)
	admin.DELETE("/:id", handler.DeleteArticle)
}