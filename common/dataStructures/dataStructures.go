package dataStructures

import (
	"github.com/gorilla/websocket"
)

type PubMessage struct {
	Topic   string `json:"topic"`
	Message []byte `json:"message"`
}

type Topic struct {
	Id         int               `json:"id"`
	Name       string            `json:"name"`
	Subscriber []*websocket.Conn `json:"subscriber"`
}

type SubcribeTo struct {
	Name string `json:"name"`
}
