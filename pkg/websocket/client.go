package websocket

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

// Client is a single websocket client
type Client struct {
	ID   string
	Conn *websocket.Conn // websocket conn for this client
	Pool *Pool // reference to Pool which manages multiple clients
	mu   sync.Mutex // mutex to protecct shared data, not used
}

type Message struct {
	Type int    `json:"type"` // type of message
	Body string `json:"body"` // content
}

func (c *Client) Read() {
	defer func() { // After reading, unregister and close the ws connection
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		msgType, p, err := c.Conn.ReadMessage() // p is actuall message in binary
		if err != nil {
			log.Println(err)
			return
		}

		msg := Message{
			Type: msgType,
			Body: string(p),
		}
		// Send msg to Broadcast channel in pool for other clients to recieve
		c.Pool.Broadcast <- msg
	}
}
