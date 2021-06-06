package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	client *Client
	data   []byte
}

type response struct {
	Method string      `json:"method"`
	Data   interface{} `json:"data"`
}

func NewResponse(method string, data interface{}) *response {
	return &response{
		Method: method,
		Data:   data,
	}
}

type errorResponse struct {
	Method string `json:"method"`
	Error  string `json:"error"`
}

func NewErrorResponse(error string) *errorResponse {
	return &errorResponse{
		Method: "error",
		Error:  error,
	}
}

func newResponse(client *Client, data interface{}) *Response {

	dataJSON, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(dataJSON))
	return &Response{
		client: client,
		data:   dataJSON,
	}
}
