package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024, // Buffer size for incoming messages
	WriteBufferSize: 1024, // Buffer size for outgoing messages
}

// Server calls this Upgrader.Upgrader method from HTTP request handler to get a *Conn
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {

	// handles CORS, true means it allows WebSocket connections from any origin.
	// not good for prod
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	// Upgrades the HTTP connection  to a WebSocket connection.
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return conn, nil
}
