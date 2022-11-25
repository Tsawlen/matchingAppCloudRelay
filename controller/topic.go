package controller

import (
	"errors"

	"github.com/Tsawlen/matchingAppCloudRelay/common/dataStructures"
	"github.com/gorilla/websocket"
)

var Topics []dataStructures.Topic

func getTopic(name string) (*dataStructures.Topic, bool) {
	for _, data := range Topics {
		if data.Name == name {
			return &data, true
		}
	}

	newTopic := dataStructures.Topic{
		Name:       name,
		Subscriber: nil,
	}

	Topics = append(Topics, newTopic)

	return &newTopic, false
}

func isSubscribed(socket *websocket.Conn, topiceName string) (*dataStructures.Topic, bool) {
	topic, exists := getTopic(topiceName)
	if !exists {
		return topic, false
	}
	for _, data := range topic.Subscriber {
		if data == socket {
			return topic, true
		}
	}
	return topic, false
}

func getSubscriberToDelete(socket *websocket.Conn, topic *dataStructures.Topic) (int, error) {
	for counter, data := range topic.Subscriber {
		if data == socket {
			return counter, nil
		}
	}
	return -1, errors.New("No Topic found with this name!")
}

func subscribeToTopic(socket *websocket.Conn, topic *dataStructures.Topic) (bool, error) {
	for counter, data := range Topics {
		if topic.Name == data.Name {
			Topics[counter].Subscriber = append(Topics[counter].Subscriber, socket)
			return true, nil
		}
	}
	return false, errors.New("Topic not found!")
}
