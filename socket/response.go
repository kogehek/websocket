package socket

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	client *Client
	data   []byte
}

type DataResponse struct {
	Method string      `json:"method"`
	Data   interface{} `json:"data"`
}

func NewDataResponse(method string, data interface{}) *DataResponse {
	return &DataResponse{
		Method: method,
		Data:   data,
	}
}

type ErrorResponse struct {
	Method string `json:"method"`
	Error  string `json:"error"`
}

func NewErrorResponse(error string) *ErrorResponse {
	return &ErrorResponse{
		Method: "error",
		Error:  error,
	}
}

func NewResponse(client *Client, data interface{}) *Response {

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
