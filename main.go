package main

import (
	"net/http"

	"github.com/1shubham7/chatterpillar/websocket"
)

func main() {
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}

func serveWS(pool *websocket.Pool, )

func setupRoutes(){
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request){
		serveWS(pool, w, r)
	})
}