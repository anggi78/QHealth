package message

import (
	"qhealth/domain"

	"github.com/labstack/echo/v4"
)

type (
	Repository interface {
		SaveMessage(msg domain.Message, recipientId string) error
		GetUnreadMessages(userID string) ([]domain.Message, error)
		IsDoctor(senderId string) (bool, error)
		GetAllMessage() ([]domain.Message, error)
	}

	Service interface {
		GetAllMessage() ([]domain.MessageResp, error)
	}

	Handler interface {
		//JoinChat(c echo.Context) error
		GetAllMessage(e echo.Context) error
	}
)