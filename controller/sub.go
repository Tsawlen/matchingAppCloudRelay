package controller

import (
	"log"
	"net/http"

	"github.com/Tsawlen/matchingAppCloudRelay/common/dataStructures"
	"github.com/gorilla/websocket"
)

var broadcaster string
var upgrader = websocket.Upgrader{
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
}

func HandleConnections(writer http.ResponseWriter, request *http.Request) {
	newWebSocket, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer newWebSocket.Close()

	for {
		var subscribeTo dataStructures.SubcribeTo
		errReadLoop := newWebSocket.ReadJSON(&subscribeTo)
		// Is subscribed?
		topic, alreadySubscribed := isSubscribed(newWebSocket, subscribeTo.Name)
		if !alreadySubscribed {
			subscribeToTopic(newWebSocket, topic)
		}
		// Delete subscription
		if errReadLoop != nil {

		}
	}
}

// Helper Functions
