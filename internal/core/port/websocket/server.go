package websocket

import (
	"context"
	"sync"
)

type WebsocketServer struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	unregister chan *Client
	register   chan *Client
	ctx        context.Context
	mu         sync.Mutex
}

func NewServer(ctx context.Context) *WebsocketServer {
	s := &WebsocketServer{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		unregister: make(chan *Client),
		register:   make(chan *Client),
		ctx:        ctx,
	}

	go s.run()

	return s
}

func (w *WebsocketServer) run() {
	for {
		select {
		case client := <-w.register:
			w.mu.Lock()
			w.clients[client] = true
			w.mu.Unlock()
		case client := <-w.unregister:
			w.mu.Lock()
			delete(w.clients, client)
			w.mu.Unlock()
		case message := <-w.broadcast:
			w.mu.Lock()
			for client := range w.clients {
				select {
				case client.send <- message:
				default:
					delete(w.clients, client)
					close(client.send)
				}
			}
			w.mu.Unlock()
		case <-w.ctx.Done():
			w.mu.Lock()
			for client := range w.clients {
				w.mu.Lock()
				close(client.send)
				w.mu.Unlock()
			}
			w.mu.Unlock()
			return
		}
	}
}

func (w *WebsocketServer) Send(client *Client, message []byte) {
	client.send <- message
}

func (w *WebsocketServer) Broadcast(message []byte) {
	w.broadcast <- message
}