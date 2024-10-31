package handler

import (
	"net/http"
	"qhealth/domain"
	"qhealth/features/queue"
	"qhealth/helpers"
	"qhealth/helpers/middleware"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type handler struct {
	serv queue.Service
}

func NewQueueService(serv queue.Service) queue.Handler {
	return &handler{
		serv: serv,
	}
}

func (h *handler) CreateQueue(e echo.Context) error {
	queueReq := domain.QueueReq{}

	if err := e.Bind(&queueReq); err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	userToken, ok := e.Get("user").(*jwt.Token)
	if !ok || !userToken.Valid {
		return helpers.CustomErr(e, "invalid token")
	}
	claims, ok := userToken.Claims.(jwt.MapClaims)
	if !ok {
		return helpers.CustomErr(e, "invalid token claims")
	}

	iduser, ok := claims["id"].(string)
	if !ok {
		return helpers.CustomErr(e, "user id not found in token")
	}

	queueReq.IdUser = iduser

	err := h.serv.CreateQueue(queueReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully create queue", nil))
}

func (h *handler) GetAllQueues(e echo.Context) error {
	queueList, err := h.serv.GetAllQueues()
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully get all data", queueList))
}

func (h *handler) GetQueueById(e echo.Context) error {
	id := e.Param("id")

	queue, err := h.serv.GetQueueByID(id)
	if err != nil{
		if err == gorm.ErrRecordNotFound {
			return helpers.CustomErr(e, err.Error())
		}
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully get data", queue))
}

func (h *handler) UpdateQueue(e echo.Context) error {
	id := e.Param("id")

	queueReq := domain.QueueReq{}
	err := e.Bind(&queueReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	err = h.serv.UpdateQueue(id, queueReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully updated data", nil))
}

func (h *handler) DeleteQueue(e echo.Context) error {
	id := e.Param("id")
	
	err := h.serv.DeleteQueue(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return helpers.CustomErr(e, err.Error())
		}
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully deleted data", nil))
}

func (h *handler) CallNextPatient(e echo.Context) error {
	queueNumber := e.Param("queue_number")

	doctorId, _, err := middleware.ExtractToken(e)  
    if err != nil {
        return helpers.CustomErr(e, "invalid token")
    }

    if queueNumber == "" {
		return helpers.CustomErr(e, "queue number is required")
	}

    err = h.serv.CallPatient(queueNumber, doctorId)
    if err != nil {
        return helpers.CustomErr(e, err.Error())
    }

    return e.JSON(http.StatusOK, helpers.SuccessResponse("patient called successfully", nil))
}
