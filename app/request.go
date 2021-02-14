package main

import (
	"encoding/json"
	"log"
)

type Request struct {
	Method string
	Body   json.RawMessage
}

func newRequest(JSON []byte) Request {
	var request Request
	err := json.Unmarshal(JSON, &request)
	if err != nil {
		log.Println(err)
	}

	return request
}
