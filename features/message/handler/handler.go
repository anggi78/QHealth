package handler

import (
	"net/http"
	"qhealth/features/message/ws"
	"qhealth/helpers/middleware"
	"strings"

	"github.com/gorilla/websocket"
)

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