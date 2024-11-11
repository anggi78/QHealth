package websocket

// type Hub struct {
// 	Clients    map[string]*Client
// 	Register   chan *Client
// 	Unregister chan *Client
// 	Broadcast  chan *Message
// }

// func NewHub() *Hub {
// 	return &Hub{
// 		Clients:    make(map[string]*Client),
// 		Register:   make(chan *Client),
// 		Unregister: make(chan *Client),
// 		Broadcast:  make(chan *Message),
// 	}
// }

// func (h *Hub) Run() {
// 	for {
// 		select {
// 		case client := <-h.Register:
// 			h.Clients[client.Id] = client
// 		case client := <-h.Unregister:
// 			if _, ok := h.Clients[client.Id]; ok {
// 				delete(h.Clients, client.Id)
// 				close(client.Message)
// 			}
// 		case message := <-h.Broadcast:
// 			for _, client := range h.Clients {
// 				if (message.IdUser != nil && client.IdUser == message.IdUser) ||
// 					(message.IdDoctor != nil && client.IdDoctor == message.IdDoctor) {
// 					client.Message <- message
// 				}
// 			}
// 		}
// 	}
// }
