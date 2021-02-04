package main

type Room struct {
	client []*Client
	id     int
}

func newRoom(client []*Client, id int) *Room {
	return &Room{
		client: client,
		id:     id,
	}
}
