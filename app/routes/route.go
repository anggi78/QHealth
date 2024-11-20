package routes

import (
	article "qhealth/features/article/router"
	view "qhealth/features/article_view/router"
	doctor "qhealth/features/doctor/router"
	notification "qhealth/features/notification/router"
	"qhealth/helpers/middleware"

	"qhealth/features/message/handler"
	"qhealth/features/message/ws"

	messages "qhealth/features/message/router"
	queue "qhealth/features/queue/router"
	status "qhealth/features/queue_status/router"
	role "qhealth/features/role/router"
	user "qhealth/features/users/router"

	//"qhealth/helpers/websocket"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Routes(e *echo.Echo, db *gorm.DB, hub *ws.Hub, validate *validator.Validate) {
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

	messageGroup := e.Group("/chat")
	messages.MessageRoute(messageGroup, db)

	message := e.Group("/msg", middleware.JwtMiddleware())
	message.GET("/ws/message", func(c echo.Context) error {
        handler.MessageHandler(hub, "", c.Response(), c.Request())
        return nil
    })

	notifictionGroup := e.Group("/notif")
	notification.NotificationRoute(notifictionGroup, db, validate)
}