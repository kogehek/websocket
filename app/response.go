package main

type Response struct {
	client *Client
	data   []byte
}

func newResponse(client *Client, data string) *Response {
	return &Response{
		client: client,
		data:   []byte(data),
	}
}
