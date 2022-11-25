package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Tsawlen/matchingAppCloudRelay/common/dataStructures"
)

func InitPub() {

}

func ReceiveMessage(writer http.ResponseWriter, request *http.Request) {

	topicIn := request.Header.Get("topic")
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

	message := dataStructures.PubMessage{
		Topic:   topicToUse.Name,
		Message: messageReceived,
	}

	json.NewEncoder(writer).Encode(message)

}
