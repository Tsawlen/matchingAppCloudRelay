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

func UnsubscribeFromTopic(socket *websocket.Conn, topic *dataStructures.Topic) (bool, error) {
	for counter, data := range Topics {
		if topic.Name == data.Name {
			for counter2, data := range data.Subscriber {
				if data == socket {
					if (len(Topics[counter].Subscriber) - 1) <= 0 {
						Topics[counter].Subscriber = []*websocket.Conn{}
					} else if (len(Topics[counter].Subscriber) - 1) <= counter2 {
						Topics[counter].Subscriber = Topics[counter].Subscriber[:len(Topics[counter].Subscriber)-1]
					} else {
						Topics[counter].Subscriber = append(Topics[counter].Subscriber[:counter2], Topics[counter].Subscriber[counter2+1])
					}
				}
			}

			return true, nil
		}
	}
	return false, errors.New("Topic not found!")
}
