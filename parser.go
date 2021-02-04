package main

import (
	"errors"
	"reflect"
)

var Methods = map[string]interface{}{
	"getName":   getName,
	"broadcast": broadcast,
}

func getName(c *Client, request Request) {
	c.hub.self <- newResponse(c, c.id)
}

func broadcast(c *Client, request Request) {
	c.hub.broadcast <- newResponse(c, request.Body)
}

func call(funcName string, params ...interface{}) (result interface{}, err error) {
	f := reflect.ValueOf(Methods[funcName])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is out of index.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	var res []reflect.Value
	res = f.Call(in)
	result = nil
	if res != nil {
		result = res[0].Interface()
	}
	return
}

func ParseJSON(JSON []byte, c *Client) {
	request := newRequest(JSON)
	call(request.Method, c, request)
}
