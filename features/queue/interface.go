package queue

import (
	"qhealth/domain"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		CreateQueue(queue domain.Queue) error
		GetAllQueues() ([]domain.Queue, error)
		GetQueueByID(id string) (*domain.Queue, error)
		GetQueueStatusByName(statusName string, status *domain.QueueStatus) error
		UpdateQueue(id string, queue domain.Queue) error
		DeleteQueue(id string) error
	}

	Service interface {
		CreateQueue(queueReq domain.QueueReq) error
		GetAllQueues() ([]domain.QueueResp, error)
		GetQueueByID(id string) (*domain.QueueResp, error)
		UpdateQueue(id string, queue domain.QueueReq) error
		DeleteQueue(id string) error
	}

	Handler interface {
		CreateQueue(e echo.Context) error
		GetAllQueues(e echo.Context) error
		GetQueueById(e echo.Context) error
		UpdateQueue(e echo.Context) error
		DeleteQueue(e echo.Context) error
	}
)