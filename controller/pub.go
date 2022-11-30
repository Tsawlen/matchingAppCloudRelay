package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/Tsawlen/matchingAppCloudRelay/common/dataStructures"
	"github.com/google/uuid"
)

func InitPub() {

}

func ReceiveMessage(writer http.ResponseWriter, request *http.Request) {

	topicIn := request.Header.Get("topic")
	service := request.Header.Get("service")
	body := request.Body
	defer body.Close()

	//Get Topic
	topicToUse, exists := getTopic(topicIn)
	if !exists {
		// Save Topic to DB
	}

	messageReceived, errRead := ioutil.ReadAll(body)
	if errRead != nil {
		log.Println(errRead)
	}

	internalMessage := buildInternalMessage(&messageReceived, topicToUse, service)

	message := dataStructures.PubMessage{
		Id:      internalMessage.Id,
		Topic:   internalMessage.Topic.Name,
		Message: internalMessage.Message,
		Service: internalMessage.Service,
	}

	json.NewEncoder(writer).Encode(message)

	Broadcaster <- internalMessage
}

// Helper

func buildInternalMessage(messageReceived *[]byte, topic *dataStructures.Topic, service string) dataStructures.InternalMessage {
	receivedAt := time.Now().Local()
	messageId, _ := uuid.NewRandom()
	msg := dataStructures.InternalMessage{
		Id:         messageId,
		Topic:      *topic,
		Message:    *messageReceived,
		Service:    service,
		ReceivedAt: receivedAt,
	}
	return msg
}
