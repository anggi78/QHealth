package handler

import (
	"net/http"
	naivebayes "qhealth/features/naive-bayes"
	"qhealth/helpers"

	"github.com/labstack/echo/v4"
)

type handler struct {
	serv naivebayes.Service
}

func NewNaiveHandler(serv naivebayes.Service) naivebayes.Handler {
	return &handler{
		serv: serv,
	}
}

func (h *handler) GetAllPatients(e echo.Context) error {
	patientList, err := h.serv.GetAllPatients()
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully get all data", patientList))
}

func (h *handler) ClassifyPatients(e echo.Context) error {
    if err := h.serv.ClassifyPatients(); err != nil {
        return helpers.CustomErr(e, err.Error())
    }

    patientList, err := h.serv.GetAllPatients()
    if err != nil {
        return helpers.CustomErr(e, err.Error())
    }

    return e.JSON(http.StatusOK, helpers.SuccessResponse("patients classified successfully", patientList))
}