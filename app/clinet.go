package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"time"
	"websocket/model"
	"websocket/store"

	"github.com/gorilla/websocket"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

type Client struct {
	id    string
	hub   *Hub
	conn  *websocket.Conn
	send  chan []byte
	store *store.Store
	user  *model.User
	auth  bool
}

func newClient(hub *Hub, conn *websocket.Conn, send chan []byte, store *store.Store) *Client {
	return &Client{
		id:    StringWithCharset(5, charset),
		hub:   hub,
		conn:  conn,
		send:  send,
		store: store,
		auth:  false,
	}
}

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func (c *Client) starAuth() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	go func() {
		time.Sleep(20 * time.Second)
		if !c.auth {
			c.hub.unregister <- c
			c.conn.Close()
			return
		}
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		fmt.Println(string(message))
		request := newRequest(message)
		call("auth", c, request)
	}
}

func (c *Client) read() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		// newMessage(message, c)
		fmt.Print(message)
		ParseJSON(message, c)
	}
}

func (c *Client) write() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write([]byte(message))

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
