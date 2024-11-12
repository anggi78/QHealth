package handler

import (
	"net/http"
	"qhealth/features/message"
	"qhealth/features/message/ws"
	"qhealth/helpers"
	"qhealth/helpers/middleware"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type handler struct {
	serv message.Service
}

func NewMessageHandler(serv message.Service) message.Handler {
	return &handler{
		serv: serv,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func MessageHandler(hub *ws.Hub, userId string, w http.ResponseWriter, r *http.Request) {
	tokenStr := r.Header.Get("Authorization")
    if tokenStr == "" {
        http.Error(w, "Authorization header missing", http.StatusUnauthorized)
        return
    }

    tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

    userId, _, err := middleware.ExtractTokenFromString(tokenStr)
    if err != nil {
        http.Error(w, "Unauthorized", http.StatusUnauthorized)
        return
    }

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Gagal meng-upgrade ke WebSocket", http.StatusInternalServerError)
		return
	}

	client := &ws.Client{
		Conn: conn,
		UserId: userId,
		Hub: hub,
	}

	hub.Register <- client

	go func() {
		defer func() {
			hub.Unregister <- client
		}()

		for {
            var message ws.Message
            err := conn.ReadJSON(&message)
            if err != nil {
                break
            }

            message.SenderId = userId
            hub.Broadcast <- message
        }
	}()
}

func (h *handler) GetAllMessage(e echo.Context) error {
	messageList, err := h.serv.GetAllMessage()
	if err != nil {
		return helpers.CustomErr(e, err.Error())
	}

	return e.JSON(http.StatusOK, helpers.SuccessResponse("successfully get all data", messageList))
}