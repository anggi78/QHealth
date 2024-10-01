package queuestatus

import (
	"qhealth/domain"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		CreateStatus(status domain.QueueStatus) error
		GetAllStatus() ([]domain.QueueStatus, error)
	}

	Service interface {
		CreateStatus(statusReq domain.QueueStatusReq) error
		GetAllStatus() ([]domain.QueueStatusResp, error)
	}

	Handler interface {
		CreateStatus(e echo.Context) error
		GetAllStatus(e echo.Context) error
	}
)