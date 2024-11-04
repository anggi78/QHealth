package queue

import (
	"qhealth/domain"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		CreateQueue(queue domain.Queue) error
		CreateOfflineQueue(queueNumber, queuePosition, statusId string) error
		GetAllQueues() ([]domain.Queue, error)
		GetQueueByID(id string) (*domain.Queue, error)
		GetQueueStatusByName(statusName string) (*domain.QueueStatus, error)
		GetLastQueue() (domain.Queue, error)
		CountWaitingQueues(doctorID, userQueue string, statusID string) (int64, error)
		UpdateQueue(id string, queue domain.Queue) error
		DeleteQueue(id string) error
		UpdateQueueStatus(queueNumber, statusID string, calledAt time.Time) error
	}

	Service interface {
		CreateQueue(queueReq domain.QueueReq) error
		GetAllQueues() ([]domain.QueueResp, error)
		GetQueueByID(id string) (*domain.QueueResp, error)
		UpdateQueue(id string, queue domain.QueueReq) error
		DeleteQueue(id string) error
		CallPatient(queueNumber, doctorID string) error
	}

	Handler interface {
		CreateQueue(e echo.Context) error
		GetAllQueues(e echo.Context) error
		GetQueueById(e echo.Context) error
		UpdateQueue(e echo.Context) error
		DeleteQueue(e echo.Context) error
		CallNextPatient(e echo.Context) error
	}
)