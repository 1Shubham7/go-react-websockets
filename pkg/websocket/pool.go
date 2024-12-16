package websocket

import "fmt"

type Pool struct {
	Register   chan *Client // Register the Client
	Unregister chan *Client // Unregister the Client
	Clients    map[*Client]bool
	Broadcast  chan Message // Broadcast the Message
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
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Println("size of connection pool: ", len(pool.Clients))
			for client := range pool.Clients {
				fmt.Println(client)
				client.Conn.WriteJSON(Message{
					Type: 1, Body: "New User Joined...",
				})
			}
			break
		case client := <- pool.Unregister:
			delete(pool.Clients, client)
			fmt.Println("size of connection pool: ", len(pool.Clients))
			for client := range pool.Clients {
				client.Conn.WriteJSON(Message{
					Type: 1, Body: "User Left...",
				})
			}
			break
		case message := <- pool.Broadcast:
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