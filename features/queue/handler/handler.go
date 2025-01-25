package handler

import (
	"net/http"
	"qhealth/domain"
	"qhealth/features/queue"
	"qhealth/helpers"
	"qhealth/helpers/middleware"
	"strconv"

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

	queueReq.IdUser = &iduser

	err := h.serv.CreateQueue(queueReq)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully create queue", nil))
}

func (h *handler) GetAllQueues(e echo.Context) error {
	page, _ := strconv.Atoi(e.QueryParam("page"))
	pageSize, _ := strconv.Atoi(e.QueryParam("pageSize"))

	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}

	queueList, totalItems, err := h.serv.GetAllQueues(page, pageSize)
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	currentPage, allPages := helpers.CalculatePaginationValues(page, pageSize, totalItems)
	pagination := helpers.PaginationResponse{
		CurrentPage: currentPage,
		NextPage:    helpers.GetNextPage(currentPage, allPages),
		PrevPage:    helpers.GetPrevPage(currentPage),
		AllPages:    allPages,
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponsePage("successfully get all data", queueList, pagination))
}

func (h *handler) GetAllQueuesAdmin(e echo.Context) error {
	queueList, _, err := h.serv.GetAllQueuesAdmin(true, 0, 0)
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

func (h *handler) CompletePatient(e echo.Context) error {
	queueNumber := e.Param("queue_number")

	doctorId, _, err := middleware.ExtractToken(e)  
    if err != nil {
        return helpers.CustomErr(e, "invalid token")
    }

    if queueNumber == "" {
		return helpers.CustomErr(e, "queue number is required")
	}

    err = h.serv.CompleteQueue(queueNumber, doctorId)
    if err != nil {
        return helpers.CustomErr(e, err.Error())
    }

    return e.JSON(http.StatusOK, helpers.SuccessResponse("patient queue completed successfully", nil))
}

func (h *handler) CancelQueuePatient(e echo.Context) error {
	queueNumber := e.Param("queue_number")

	doctorId, _, err := middleware.ExtractToken(e)  
    if err != nil {
        return helpers.CustomErr(e, "invalid token")
    }

    if queueNumber == "" {
		return helpers.CustomErr(e, "queue number is required")
	}

    err = h.serv.CancelQueue(queueNumber, doctorId)
    if err != nil {
        return helpers.CustomErr(e, err.Error())
    }

    return e.JSON(http.StatusOK, helpers.SuccessResponse("patient queue cancelled successfully", nil))
}