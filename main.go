package main

import (
	"log"
	"net/http"
	"websocket/auth"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := newClient(hub, conn, make(chan []byte))
	client.hub.register <- client

	go client.write()
	go client.read()
}

func main() {

	// auth.Time(auth.ExampleNewWithClaims_standardClaims())
	auth.Time("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJiYXIiLCJuYW1lIjoiSm9obiBEb2UiLCJpYXQiOiIxNTE2MjM5MDIyIiwiZXhwIjoxNjEyNDQ1NzU1LCJpc3MiOiJ0ZXN0In0.5uj1kF_DSgWsdqGmZk-iwzzSwxskfQMMATN_xGhmDZo")
	// hub := newHub()
	// go hub.run()
	// http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
	// 	serveWs(hub, w, r)
	// })
	// log.Fatal(http.ListenAndServe(":8080", nil))
}
