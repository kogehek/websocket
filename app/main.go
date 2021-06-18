package main

import (
	"log"
	"net/http"
	"websocket/socket"
	"websocket/store"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWs(hub *socket.Hub, store *store.Store, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := socket.NewClient(hub, conn, make(chan []byte), store)
	client.Hub.Register <- client

	// client.starAuth()
	go client.Write()
	go client.Read()
}

func main() {
	store := store.New("host=localhost port=5432 dbname=websocket sslmode=disable user=root password=root")
	defer store.DB.Close()

	hub := socket.NewHub()
	go hub.Run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, store, w, r)
	})
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", nil))
}
