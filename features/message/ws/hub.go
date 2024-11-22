package ws

import (
	"log"
	"qhealth/domain"
	"qhealth/features/doctor"
	"qhealth/features/message"
	"qhealth/features/users"
	"qhealth/helpers"
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
	Clients          map[string]*Client
	Broadcast        chan Message
	Register         chan *Client
	Unregister       chan *Client
	Repository       message.Repository
	RepositoryUser   users.Repository
	RepositoryDoctor doctor.Repository
	mu               sync.Mutex
}

type Message struct {
	SenderId   string `json:"sender_id"`
	ReceiverId string `json:"receiver_id"`
	Body       string `json:"body"`
}

func NewHub(repo message.Repository, repoUser users.Repository, repoDoctor doctor.Repository) *Hub {
	return &Hub{
		Clients:          make(map[string]*Client),
		Broadcast:        make(chan Message),
		Register:         make(chan *Client),
		Unregister:       make(chan *Client),
		Repository:       repo,
		RepositoryUser:   repoUser,
		RepositoryDoctor: repoDoctor,
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
			log.Printf("Pesan diterima: %+v", message)
			msg := domain.Message{
				MessageBody: message.Body,
				IdUser:      message.SenderId,
				CreateDate:  time.Now(),
			}

			if message.SenderId == "" || message.ReceiverId == "" {
				log.Printf("Gagal: SenderId atau ReceiverId kosong: SenderId=%s, ReceiverId=%s", message.SenderId, message.ReceiverId)
				continue
			}

			isDoc, err := hub.Repository.IsDoctor(message.SenderId)
			if err != nil {
				log.Printf("Gagal memeriksa peran pengirim: %v", err)
				continue
			}

			if isDoc {
				msg.IdDoctor = message.SenderId
				msg.IdUser = message.ReceiverId
			} else {
				msg.IdUser = message.SenderId
				msg.IdDoctor = message.ReceiverId
			}

			if msg.IdUser == "" || msg.IdDoctor == "" {
				log.Printf("Pengaturan ID gagal: IdUser atau IdDoctor kosong")
				continue
			}

			if err := hub.Repository.SaveMessage(msg, message.ReceiverId); err != nil {
				log.Printf("Failed to save message: %v", err)
				continue
			}

			hub.mu.Lock()
			recipient, ok := hub.Clients[message.ReceiverId]

			email := ""

			emailUser, errUser := hub.RepositoryUser.FindById(message.ReceiverId)
			if errUser == nil {
				email = emailUser.Email
				log.Printf("Email found in User repository for ReceiverId %s: %s", message.ReceiverId, email)
			} else {
				emailDoctor, errDoctor := hub.RepositoryDoctor.FindById(message.ReceiverId)
				if errDoctor == nil {
					email = emailDoctor.Email
					log.Printf("Email found in Doctor repository for ReceiverId %s: %s", message.ReceiverId, email)
				} else {
					log.Printf("Failed to fetch email for ReceiverId %s: UserError=%v, DoctorError=%v", message.ReceiverId, errUser, errDoctor)
					return
				}
			}

			if ok {
				if err := recipient.Conn.WriteJSON(message); err != nil {
					log.Printf("Error sending message to %s: %v", message.ReceiverId, err)
				} else {
					log.Printf("Message sent from %s to %s: %s", message.SenderId, message.ReceiverId, message.Body)
				}
			} else {
				log.Printf("Recipient %s is not connected. Sending email notification.", message.ReceiverId)
				if err := helpers.SendEmailNotification(email); err != nil {
					log.Printf("Failed to send email notification to %s: %v", email, err)
				} else {
					log.Printf("Email notification sent to %s successfully.", email)
				}
			}

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
