package handler

import (
	"encoding/json"
	"fmt"
	"qhealth/domain"
	"qhealth/features/notification"
	"time"

	"github.com/labstack/echo/v4"
)

type handler struct {
	serv notification.Service
}

func NewNotificationHandler(serv notification.Service) notification.Handler {
	return &handler{
		serv: serv,
	}
}

func (h *handler) FindAllNotification(c echo.Context) error {
	c.Response().Header().Set("Content-Type", "text/event-stream")
	c.Response().Header().Set("Cache-Control", "no-cache")
	c.Response().Header().Set("Connection", "keep-alive")

	var lastUpdate time.Time

	messageChan := make(chan string)

	for {
		select {
		case <- c.Request().Context().Done():
			close(messageChan)
			return nil
		default:
			result, _ := h.serv.FindAllNotification(c)
			if len(result) == 0 {
				message := fmt.Sprintf("data: %s\n\n", "null")
				fmt.Fprintf(c.Response(), message)
				c.Response().Flush()
				lastUpdate = time.Time{}
			}
			if len(result) > 0 && result[0].CreatedAt != lastUpdate {
				results := domain.ListNotificationToResp(result)
				data, _ := json.Marshal(results)
				message := fmt.Sprintf("data: %s\n\n", data)
				fmt.Fprintf(c.Response(), message)
				c.Response().Flush()
				lastUpdate = result[0].CreatedAt
			}
		}
		time.Sleep(2 * time.Second)
	}
}