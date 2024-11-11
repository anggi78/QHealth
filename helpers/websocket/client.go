package websocket

// import (
// 	configs "qhealth/app/drivers"
// 	"qhealth/domain"
// 	"time"

// 	"github.com/gorilla/websocket"
// )

// type Client struct {
// 	Conn     *websocket.Conn
// 	Message  chan *Message
// 	Id       string
// 	IdUser   *string
// 	IdDoctor *string
// }

// type Message struct {
// 	IdUser   *string `json:"id_user"`
// 	IdDoctor *string `json:"id_doctor"`
// 	Body     string  `json:"body"`
// 	IsRead   bool    `json:"is_read"`
// }

// func (c *Client) WriteMessage() {
// 	defer func() {
// 		c.Conn.Close()
// 	}()

// 	for {
// 		message, ok := <-c.Message
// 		if !ok {
// 			c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
// 			return
// 		}
// 		c.Conn.WriteJSON(message)

// 	}
// }

// func (c *Client) ReadMessage(hub *Hub) {
// 	defer func() {
// 		hub.Unregister <- c
// 		c.Conn.Close()
// 	}()

// 	for {
// 		_, msg, err := c.Conn.ReadMessage()
// 		if err != nil {
// 			break
// 		}

// 		message := &Message{
// 			IdUser:   c.IdUser,
// 			IdDoctor: c.IdDoctor,
// 			Body:     string(msg),
// 			IsRead:   false,
// 		}

// 		dbMessage := domain.Message{
// 			Body:       message.Body,
// 			IdUser:     *message.IdUser,
// 			IdDoctor:   *message.IdDoctor,
// 			CreateDate: time.Now(),
// 		}
// 		configs.InitDB().Create(&dbMessage)

// 		hub.Broadcast <- message
// 	}
// }
