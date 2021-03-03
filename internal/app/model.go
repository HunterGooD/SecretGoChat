package app

import (
	"crypto/rsa"

	"github.com/gorilla/websocket"
)

// Room структура комнаты
type Room struct {
	Users         []User
	ServerPrivate *rsa.PrivateKey
	AESKey        string
}

// User описывает структуру пользователя
type User struct {
	UUID string
	Conn *websocket.Conn
}

// Package пакет обмен данными
type Package struct {
	Head HeadPackage `json:"head"`
	Body BodyPackage `json:"body"`
}

// HeadPackage заголовки отправленных данных
type HeadPackage struct {
	Rand       string `json:"rand"`
	Title      string `json:"title"`
	Sender     string `json:"sender"`
	SessionKey string `json:"session"`
}

//BodyPackage Тело отправленных данных
type BodyPackage struct {
	Data string `json:"data"`
	Hash string `json:"hash"`
	Sign string `json:"sign"`
}
