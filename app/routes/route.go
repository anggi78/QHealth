package routes

import (
	"qhealth/features/users/router"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	userGroup := e.Group("/users")
	router.UserRoute(userGroup, db)
}