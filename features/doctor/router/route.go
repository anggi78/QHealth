package router

import (
	"qhealth/features/doctor/handler"
	"qhealth/features/doctor/repository"
	"qhealth/features/doctor/service"
	"qhealth/helpers/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DoctorRoute(d *echo.Group, db *gorm.DB) {
	repo := repository.NewDoctorRepository(db)
	serv := service.NewDoctorService(repo)
	handler := handler.NewDoctorHandler(serv)

	d.POST("/register", handler.Register)
	d.POST("/login", handler.Login)
	d.POST("/forgot-password", handler.ForgotPassword)

	auth := d.Group("/profile", middleware.JwtMiddleware())
	auth.POST("/forgot", handler.ChangePassForgot)
	auth.POST("/change", handler.ChangePass)
	auth.PUT("", handler.UpdateProfile)
	auth.DELETE("", handler.DeleteProfile)
}