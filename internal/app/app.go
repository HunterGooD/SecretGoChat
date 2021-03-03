package app

import "github.com/gorilla/websocket"

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

// App главная структура приложения
type App struct {
	Rooms []Room
}
