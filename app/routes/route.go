package routes

import (
	article "qhealth/features/article/router"
	user "qhealth/features/users/router"
	role "qhealth/features/role/router"
	view "qhealth/features/article_view/router"
	doctor "qhealth/features/doctor/router"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB) {
	userGroup := e.Group("/users")
	user.UserRoute(userGroup, db)
	
	articleGroup := e.Group("/article")
	article.ArticleRoute(articleGroup, db)

	roleGroup := e.Group("/role")
	role.RoleRoute(roleGroup, db)

	viewGroup := e.Group("/view")
	view.ViewRoute(viewGroup, db)

	doctroGroup := e.Group("/doctor")
	doctor.DoctorRoute(doctroGroup, db)
}