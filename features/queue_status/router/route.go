package router

import (
	"qhealth/features/queue_status/handler"
	"qhealth/features/queue_status/repository"
	"qhealth/features/queue_status/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func StatusRoute(s *echo.Group, db *gorm.DB) {
	repo := repository.NewStatusRepository(db)
	serv := service.NewStatusService(repo)
	handler := handler.NewStatusHandler(serv)

	s.POST("", handler.CreateStatus)
	s.GET("", handler.GetAllStatus)
}