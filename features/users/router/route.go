package router

import (
	"qhealth/features/users/handler"
	"qhealth/features/users/repository"
	"qhealth/features/users/service"
	"qhealth/helpers/middleware"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserRoute(u *echo.Group, db *gorm.DB) {
	repo := repository.NewUserRepository(db)
	serv := service.NewService(repo)
	handler := handler.NewUserHandler(serv)

	u.POST("/register", handler.Register)
	u.POST("/register-admin", handler.RegisterAdmin)
	u.POST("/login", handler.Login)
	u.POST("/forgot-password", handler.ForgotPassword)
	u.GET("/initialize", handler.InitializeRolesAndPermissions)
	
	auth := u.Group("/profile", middleware.JwtMiddleware())
	mw := middleware.NewMiddleware(db)
	auth.POST("/forgot", handler.ChangePassForgot)
	auth.POST("/change", handler.ChangePass)
	auth.PUT("", handler.UpdateUser, mw.Authorize("edit"))
	auth.DELETE("", handler.DeleteUser, mw.Authorize("delete"))
}