package websocket

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Server calls this Upgrader.Upgrader method from HTTP request handler to get a *Conn
func Upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return conn, nil
}
