package main

import (
	"fmt"
	"net/http"

	"github.com/1shubham7/chatterpillar/pkg/websocket"
)

func main() {
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Printf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	// So we register the client and then read what's coming from the client, then unregister it and close
	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}
