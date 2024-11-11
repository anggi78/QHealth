package ws

import (
	"log"
	"qhealth/domain"
	"qhealth/features/message"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	Conn   *websocket.Conn
	UserId string
	Hub    *Hub
}

type Hub struct {
	Clients    map[string]*Client
	Broadcast  chan Message
	Register   chan *Client
	Unregister chan *Client
	Repository message.Repository
	mu         sync.Mutex
}

type Message struct {
	SenderId   string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
	Body       string `json:"body"`
}

func NewHub(repo message.Repository) *Hub {
	return &Hub{
		Clients:    make(map[string]*Client),
		Broadcast:  make(chan Message),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Repository: repo,
	}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.Register:
			hub.mu.Lock()
			hub.Clients[client.UserId] = client
			log.Printf("User %s connected.", client.UserId)

			undeliveredMessages, err := hub.Repository.GetUnreadMessages(client.UserId)
			if err == nil {
				for _, msg := range undeliveredMessages {
					client.Conn.WriteJSON(msg)
				}
			} else {
				log.Printf("Failed to retrieve undelivered messages for %s: %v", client.UserId, err)
			}

			hub.mu.Unlock()

		case client := <-hub.Unregister:
			hub.mu.Lock()
			delete(hub.Clients, client.UserId)
			client.Conn.Close()
			hub.mu.Unlock()

		case message := <-hub.Broadcast:
			msg := domain.Message{
				MessageBody: message.Body,
				IdUser: message.SenderId,
				CreateDate: time.Now(),
			}

			if message.ReceiverId != "" {
				msg.IdDoctor = message.ReceiverId
			}

			if err := hub.Repository.SaveMessage(msg, message.ReceiverId); err != nil {
				log.Printf("Failed to save message: %v", err)
				continue
			}

			hub.mu.Lock()
			recipient, ok := hub.Clients[message.ReceiverId]
			hub.mu.Unlock()

			if ok {
				if err := recipient.Conn.WriteJSON(message); err != nil {
					log.Printf("Error sending message to %s: %v", message.ReceiverId, err)
				} else {
					log.Printf("Message from %s to %s: %s", message.SenderId, message.ReceiverId, message.Body)
				}
			} else {
				log.Printf("Failed to send message, recipient %s not connected.", message.ReceiverId)
			}
		}
	}
}
