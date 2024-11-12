package router

import (
	"qhealth/features/message/handler"
	"qhealth/features/message/repository"
	"qhealth/features/message/service"
	"qhealth/helpers/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func MessageRoute(m *echo.Group, db *gorm.DB) {
	repo := repository.NewMessageRepository(db)
	serv := service.NewMessageService(repo)
	handler := handler.NewMessageHandler(serv)

	messages := m.Group("/message", middleware.JwtMiddleware())
	messages.GET("", handler.GetAllMessage)
}