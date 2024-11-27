package notification

import (
	"qhealth/domain"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		FindAll() ([]domain.Notification, error)
		SaveNotification(notification domain.Notification) error
	}

	Service interface {
		FindAllNotification(c echo.Context) ([]domain.Notification, error)
	}

	Handler interface {
		FindAllNotification(c echo.Context) error
	}
)