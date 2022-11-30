package dataStructures

import (
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type PubMessage struct {
	Id      uuid.UUID `json:"id"`
	Topic   string    `json:"topic"`
	Message []byte    `json:"message"`
	Service string    `json:"service"`
}

type Topic struct {
	Id         int               `json:"id"`
	Name       string            `json:"name"`
	Subscriber []*websocket.Conn `json:"subscriber"`
}

type InternalMessage struct {
	Id         uuid.UUID
	Topic      Topic
	Message    []byte
	Service    string
	ReceivedAt time.Time
}

type OutgoingMessage struct {
	Id          uuid.UUID `json:"id"`
	Topic       string    `json:"topic"`
	Message     []byte    `json:"message"`
	Service     string    `json:"service"`
	ReceivedAt  time.Time `json:"receivedAt"`
	DeliveredAt time.Time `json:"deliveredAt"`
}

type SubcribeTo struct {
	Name string `json:"name"`
}
