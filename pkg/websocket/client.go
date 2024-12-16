package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID string
	Conn *websocket.Conn
	Pool *Pool
	mu sync.Mutex
}

func (c *Client) Read() {
	defer func() { // After reading, unregister and close the ws connection
		c.Pool.Unregister <- c
		c.Conn.Close()
	} ()

	
}