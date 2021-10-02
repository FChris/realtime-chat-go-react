package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fchris/realtime-chat-go-react/pkg/websocket"
)

func serveWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r)

	if err != nil {
		log.Println(err)
	}

	go websocket.Writer(ws)
	websocket.Reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/ws", serveWebSocket)
}

func main() {
	fmt.Println("Distributed Chat App v0.0.1")
	setupRoutes()
	http.ListenAndServe(":8080", nil)

}
