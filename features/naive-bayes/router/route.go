package router

import (
	"qhealth/features/naive-bayes/handler"
	"qhealth/features/naive-bayes/repository"
	"qhealth/features/naive-bayes/service"
	"qhealth/helpers/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func PatientRoute(p *echo.Group, db *gorm.DB) {
	repo := repository.NewNaiveRepository(db)
	serv := service.NewNaiveService(repo)
	handler := handler.NewNaiveHandler(serv)

	patient := p.Group("/patient", middleware.JwtMiddleware())
	patient.GET("", handler.GetAllPatients)
}