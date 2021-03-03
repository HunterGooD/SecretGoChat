package main

import "crypto/rsa"

var users = make(map[string]string)

// User cструктура пользователя отправляющего сообщения
type User struct {
	Private *rsa.PrivateKey
}
