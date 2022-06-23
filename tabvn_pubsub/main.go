package main

import (
	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func autoId() (string) {
	return (go.uuid).Must((go.uuid).NewV4()).String()
}