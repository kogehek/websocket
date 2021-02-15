package main

import (
	"log"
	"net/http"

	// "websocket/auth"
	// "websocket/auth"

	"websocket/store"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func serveWs(hub *Hub, store *store.Store, w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := newClient(hub, conn, make(chan []byte), store)
	client.hub.register <- client

	go client.write()
	go client.read()
}

func main() {
	// fmt.Println(auth.ExampleNewWithClaims_standardClaims())
	// auth.Time("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJiYXIiLCJuYW1lIjoiSm9obiBEb2UiLCJpYXQiOiIxNTE2MjM5MDIyIiwiZXhwIjoxNjEzNDA0MDIzLCJpc3MiOiJ0ZXN0In0.5So2-YcZHUxHZENG7Y0Hy5TzQCIqlSJ-jA0CZtaOct0")
	store := store.New("host=localhost port=5432 dbname=websocket sslmode=disable user=root password=root")
	defer store.DB.Close()

	hub := newHub()
	go hub.run()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, store, w, r)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
