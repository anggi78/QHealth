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

	article := a.Group("/articles", middleware.JwtMiddleware())
	mw := middleware.NewMiddleware(db)
	
	// admin
	article.POST("", handler.CreateArticle, mw.AuthorizeAdmin())
	article.PUT("/:id", handler.UpdateArticle, mw.AuthorizeAdmin())
	article.DELETE("/:id", handler.DeleteArticle, mw.AuthorizeAdmin())

	// user
	article.GET("", handler.GetAllArticle, mw.Authorize("read"))
	article.GET("/:id", handler.GetArticleById, mw.Authorize("read"))
	article.GET("/latest", handler.GetLatestArticle, mw.Authorize("read"))
}