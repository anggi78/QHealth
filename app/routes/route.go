package routes

import (
	article "qhealth/features/article/router"
	user "qhealth/features/users/router"
	role "qhealth/features/role/router"
	view "qhealth/features/article_view/router"
	doctor "qhealth/features/doctor/router"
	status "qhealth/features/queue_status/router"
	queue "qhealth/features/queue/router"

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

	statusGroup := e.Group("/status")
	status.StatusRoute(statusGroup, db)

	queueGroup := e.Group("/queues")
	queue.QueueRoute(queueGroup, db)
}