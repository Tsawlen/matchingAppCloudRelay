package main

import (
	"log"
	"net/http"

	"github.com/Tsawlen/matchingAppCloudRelay/controller"
	"github.com/gorilla/mux"
)

func main() {

	controller.InitPub()

	go controller.HandleIncoming()

	router := mux.NewRouter()

	router.HandleFunc("/publish", controller.ReceiveMessage).Methods("PUT")
	router.HandleFunc("/subscribe", controller.HandleConnections)

	server := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}
	log.Println("CloudRelay is started and listens to 0.0.0.0:8082")
	server.ListenAndServe()
}
