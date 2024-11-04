package handler

import (
	"net/http"
	"qhealth/domain"
	queuestatus "qhealth/features/queue_status"
	"qhealth/helpers"

	"github.com/labstack/echo/v4"
)

type handler struct {
	serv queuestatus.Service
}

func NewStatusHandler(serv queuestatus.Service) queuestatus.Handler {
	return &handler{
		serv: serv,
	}
}

func (h *handler) CreateStatus(e echo.Context) error {
	statusReq := domain.QueueStatusReq{}

	if err := e.Bind(&statusReq); err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	err := h.serv.CreateStatus(statusReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully create status", nil))
}

func (h *handler) GetAllStatus(e echo.Context) error {
	statusList, err := h.serv.GetAllStatus()
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully get all data", statusList))
}