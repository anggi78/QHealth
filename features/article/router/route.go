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
	mw := middleware.NewMiddleware(db)
	
	// admin
	admin.POST("", handler.CreateArticle, mw.AuthorizeAdmin())
	admin.PUT("/:id", handler.UpdateArticle, mw.AuthorizeAdmin())
	admin.DELETE("/:id", handler.DeleteArticle, mw.AuthorizeAdmin())

	// user
	admin.GET("", handler.GetAllArticle, mw.Authorize("read"))
	admin.GET("/:id", handler.GetArticleById, mw.Authorize("read"))
}