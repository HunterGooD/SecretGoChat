package main

import (
	"crypto/rsa"
	"syscall/js"
)

var users = make(map[string]string)

// User cструктура пользователя отправляющего сообщения
type User struct {
	Private *rsa.PrivateKey
}

func checkAndCreateRoom() js.Func {
	return js.FuncOf(func(this js.Value, input []js.Value) interface{} {
		return ""
	})
}

func main() {
	js.Global().Set("checkAndCreateRoom", checkAndCreateRoom())
	<-make(chan bool)
}
