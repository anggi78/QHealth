package routes

import (
	article "qhealth/features/article/router"
	user "qhealth/features/users/router"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	userGroup := e.Group("/users")
	user.UserRoute(userGroup, db)
	
	articleGroup := e.Group("/article")
	article.ArticleRoute(articleGroup, db)
}