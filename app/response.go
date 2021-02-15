package main

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	client *Client
	data   []byte
}

type tokenResponse struct {
	Method string `json:"method"`
	Token  string `json:"token"`
}

func NewTokenResponse(token string) *tokenResponse {
	return &tokenResponse{
		Method: "token",
		Token:  token,
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

	return &Response{
		client: client,
		data:   dataJSON,
	}
}
