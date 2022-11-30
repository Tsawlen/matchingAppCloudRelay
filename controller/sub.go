package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/Tsawlen/matchingAppCloudRelay/common/dataStructures"
	"github.com/gorilla/websocket"
)

var Broadcaster = make(chan dataStructures.InternalMessage)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}

func HandleConnections(writer http.ResponseWriter, request *http.Request) {
	newWebSocket, err := upgrader.Upgrade(writer, request, nil)
	// Get Topic
	topicName := request.Header.Get("topic")
	// Check if subscribed to topic
	topic, alreadySubscribed := isSubscribed(newWebSocket, topicName)
	if !alreadySubscribed {
		// Subscribe to Topic
		subscribeToTopic(newWebSocket, topic)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer newWebSocket.Close()

	for {
		var subscribeTo dataStructures.SubcribeTo
		errReadLoop := newWebSocket.ReadJSON(&subscribeTo)
		// Delete subscription
		if errReadLoop != nil {
			UnsubscribeFromTopic(newWebSocket, topic)
			newWebSocket.Close()
			break
		}
	}
}

func HandleIncoming() {
	for {
		msg := <-Broadcaster
		log.Println("Will now deliver: " + msg.Id.String())
		DeliverToSubscribers(&msg.Topic, buildOutgoingMessage(msg))
	}
}

func DeliverToSubscribers(topic *dataStructures.Topic, message *dataStructures.OutgoingMessage) {
	for _, receiver := range topic.Subscriber {
		err := receiver.WriteJSON(message)
		if err != nil && websocket.IsCloseError(err, websocket.CloseGoingAway) {
			receiver.Close()
			UnsubscribeFromTopic(receiver, topic)
		}
	}
	log.Println("Have now delivered: " + message.Id.String())
}

// Helper Functions
func buildOutgoingMessage(message dataStructures.InternalMessage) *dataStructures.OutgoingMessage {
	deliveredAt := time.Now().Local()
	messageToReturn := dataStructures.OutgoingMessage{
		Id:          message.Id,
		Topic:       message.Topic.Name,
		Message:     message.Message,
		Service:     message.Service,
		ReceivedAt:  message.ReceivedAt,
		DeliveredAt: deliveredAt,
	}
	return &messageToReturn
}
