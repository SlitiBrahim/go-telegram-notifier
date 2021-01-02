package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"go-telegram-notifier/src/config"
	"go-telegram-notifier/src/helper"
	"log"
	"net/http"
)

func sendNotificationHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello you")
	helper.FailOnError(err)
}

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/send-notification", sendNotificationHandler).Methods("POST")

	log.Printf("Listening on localhost:%v\n", config.Config["APP_PORT"])
	err := http.ListenAndServe(fmt.Sprintf(":%v", config.Config["APP_PORT"]), router)
	helper.FailOnError(err)
}