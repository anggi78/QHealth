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
		GetAllQueues(offset, limit int) ([]domain.Queue, int, error)
		GetQueueByID(id string) (*domain.Queue, error)
		GetQueueStatusByName(statusName string) (*domain.QueueStatus, error)
		GetLastQueue(queueType string) (*domain.Queue, error)
		CountWaitingQueues(doctorID, userQueue, statusID string) (int64, error)
		CountWaitingQueuesBeforePage(queueNumber string, statusID string) (int64, error)
		UpdateQueue(id string, queue domain.Queue) error
		DeleteQueue(id string) error
		UpdateQueueStatus(queueNumber, statusID string, calledAt time.Time) error
		UpdateQueuePosition(Id, newPosition string) error
	}

	Service interface {
		CreateQueue(queueReq domain.QueueReq) error
		GetAllQueues(page, pageSize int) ([]domain.QueueResp, int, error)
		GetAllQueuesAdmin(admin bool, page, pageSize int) ([]domain.QueueResp, int, error)
		GetQueueByID(id string) (*domain.QueueResp, error)
		UpdateQueue(id string, queue domain.QueueReq) error
		DeleteQueue(id string) error
		CallPatient(queueNumber, doctorID string) error
		CompleteQueue(queueNumber, doctorID string) error
		CancelQueue(queueNumber, doctorID string) error
	}

	Handler interface {
		CreateQueue(e echo.Context) error
		GetAllQueues(e echo.Context) error
		GetAllQueuesAdmin(e echo.Context) error
		GetQueueById(e echo.Context) error
		UpdateQueue(e echo.Context) error
		DeleteQueue(e echo.Context) error
		CallNextPatient(e echo.Context) error
		CompletePatient(e echo.Context) error
		CancelQueuePatient(e echo.Context) error
	}
)