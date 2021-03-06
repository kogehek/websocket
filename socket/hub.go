package socket

type Hub struct {
	broadcast  chan *Response
	self       chan *Response
	Register   chan *Client
	unregister chan *Client
	clients    map[*Client]bool
}

func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan *Response),
		self:       make(chan *Response),
		Register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.self:
			select {
			case message.client.send <- []byte(message.data):
			default:
				close(message.client.send)
				delete(h.clients, message.client)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				if client != message.client {
					select {
					case client.send <- message.data:
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}
