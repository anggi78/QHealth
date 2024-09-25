package handler

import (
	"net/http"
	"qhealth/domain"
	"qhealth/features/role"
	"qhealth/helpers"

	"github.com/labstack/echo/v4"
)

type handler struct {
	serv role.Service
}

func NewRoleHandler(serv role.Service) role.Handler {
	return &handler{serv: serv}
}

func (h *handler) CreateRole(e echo.Context) error {
	roleReq := domain.RoleReq{}

	if err := e.Bind(&roleReq); err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	err := h.serv.CreateRole(roleReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully create role", nil))
}

func (h *handler) GetAllRole(e echo.Context) error {
	roleList, err := h.serv.GetAllRole()
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully get all data", roleList))
}