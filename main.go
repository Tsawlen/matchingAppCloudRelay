package main

import (
	"net/http"

	"github.com/Tsawlen/matchingAppCloudRelay/controller"
	"github.com/gorilla/mux"
)

func main() {

	controller.InitPub()

	router := mux.NewRouter()

	router.HandleFunc("/publish", controller.ReceiveMessage).Methods("PUT")
	router.HandleFunc("/subscribe", controller.HandleConnections)

	server := &http.Server{
		Addr:    ":8082",
		Handler: router,
	}

	server.ListenAndServe()

}
