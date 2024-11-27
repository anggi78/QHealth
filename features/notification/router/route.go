package router

import (
	"qhealth/features/notification/handler"
	"qhealth/features/notification/repository"
	"qhealth/features/notification/service"
	"qhealth/helpers/middleware"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NotificationRoute(n *echo.Group, db *gorm.DB, validate *validator.Validate) {
	repo := repository.NewNotificationRepository(db)
	serv := service.NewNotificationService(repo, validate)
	handler := handler.NewNotificationHandler(serv)

	notification := n.Group("/notification", middleware.JwtMiddleware())
	notification.GET("", handler.FindAllNotification)
}