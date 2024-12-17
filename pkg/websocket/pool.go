package websocket

import "fmt"

type Pool struct {
	Register   chan *Client     // Register the Client
	Unregister chan *Client     // Unregister the Client
	Clients    map[*Client]bool // keeps track of connected clients
	Broadcast  chan Message     // Broadcast the Message
	// these channels are like, Register is a channel where new
	// clients are sent to be registered.
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client), // just initialize the channels
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
	}
}

func (pool *Pool) Start() {
	for { // infinite loop to keep pool active
		select {
		case client := <-pool.Register: // When the client sends itself to registration channel
			pool.Clients[client] = true // add client to map
			fmt.Println("size of connection pool: ", len(pool.Clients))
			for client := range pool.Clients { // Sending to all clients
				fmt.Println(client)
				client.Conn.WriteJSON(Message{
					Type: 1, Body: "New User Joined...",
				})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client) // remove client from map
			fmt.Println("size of connection pool: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{
					Type: 1, Body: "User Left...",
				})
			}
			break
		case message := <-pool.Broadcast: // when message is sent to Broadcast channel
			for client := range pool.Clients {
				err := client.Conn.WriteJSON(message)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
