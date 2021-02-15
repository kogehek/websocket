package main

import (
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"websocket/jwt"
	"websocket/model"
)

var Methods = map[string]interface{}{
	"getName":      getName,
	"broadcast":    broadcast,
	"registration": registration,
	"auth":         auth,
}

func getName(c *Client, request Request) {
	c.hub.self <- newResponse(c, c.id)
}

func broadcast(c *Client, request Request) {
	// c.hub.broadcast <- newResponse(c, request.Body)
}

func auth(c *Client, request Request) {
	var user *model.User
	err := json.Unmarshal(request.Body, &user)
	if err != nil {
		log.Println(err)
	}
	user, err = c.store.UserRepository.Auth(user)
	if err != nil {
		c.hub.self <- newResponse(c, err.Error())
		return
	}
	token := jwt.CrateToken(user.ID)
	c.hub.self <- newResponse(c, token)
}

func registration(c *Client, request Request) {
	var user *model.User
	err := json.Unmarshal(request.Body, &user)
	if err != nil {
		log.Println(err)
	}
	err = c.store.UserRepository.Create(user)
	if err != nil {
		c.hub.self <- newResponse(c, err.Error())
		return
	}
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
