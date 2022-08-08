package main

import "fmt"

type Hub struct {
	clients    map[*Client]bool
	registy    chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{

		clients:    make(map[*Client]bool),
		registy:    make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
	}

}

func (h *Hub) run() {

	select {
	case client := <-h.registy:
		h.clients[client] = true

	case client := <-h.unregister:
		if !h.clients[client] {

			delete(h.clients, client)
			close(client.send)
		}

	case message := <-h.broadcast:

		fmt.Println(message)
		for client := range h.clients {
			fmt.Println("client")

			fmt.Println(client)
			select {
			case client.send <- message:
			default:
				close(client.send)
				delete(h.clients, client)
			}
		}

	}
	fmt.Println("huuuu...")

}
