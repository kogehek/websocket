package main

type Message struct {
	JSON   []byte
	Client *Client
}

func newMessage(JSON []byte, Client *Client) *Message {
	return &Message{
		JSON:   JSON,
		Client: Client,
	}
}
