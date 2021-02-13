package main

import (
	"errors"
	"reflect"
	"websocket/model"
)

var Methods = map[string]interface{}{
	"getName":   getName,
	"broadcast": broadcast,
	"auth":      auth,
}

func getName(c *Client, request Request) {
	c.hub.self <- newResponse(c, c.id)
}

func broadcast(c *Client, request Request) {
	c.hub.broadcast <- newResponse(c, request.Body)
}

func auth(c *Client, request Request) {
	user, err := model.NewUser("sadasdssad.dd", "3")
	if err != nil {
		c.hub.self <- newResponse(c, err.Error())
		return
	}
	c.store.UserRepository.Create(user)
	c.hub.self <- newResponse(c, "USER CREATE")
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
