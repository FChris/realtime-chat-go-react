package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fchris/realtime-chat-go-react/pkg/websocket"
)

func serveWebSocket(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		log.Println(err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWebSocket(pool, w, r)
	})
}

func main() {
	fmt.Println("Distributed Chat App v0.0.1")
	setupRoutes()
	http.ListenAndServe(":8080", nil)

}
