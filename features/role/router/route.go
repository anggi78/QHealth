package router

import (
	"qhealth/features/role/handler"
	"qhealth/features/role/repository"
	"qhealth/features/role/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RoleRoute(r *echo.Group, db *gorm.DB) {
	repo := repository.NewRoleRepository(db)
	serv := service.NewRoleService(repo)
	handler := handler.NewRoleHandler(serv)

	r.POST("", handler.CreateRole)
	r.GET("", handler.GetAllRole)
}