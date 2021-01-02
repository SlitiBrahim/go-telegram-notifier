package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"go-telegram-notifier/src/config"
	"go-telegram-notifier/src/helper"
	"log"
	"net/http"
)

func sendNotificationHandler(w http.ResponseWriter, r *http.Request) {
	var msg Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		helper.SendApiError(w, errors.New("invalid request body: cannot parse body to Message object"), http.StatusBadRequest)
		return
	}

	if msg.Message == "" {
		helper.SendApiError(w, errors.New("empty message passed"), http.StatusBadRequest)
		return
	}

	_, err = fmt.Fprintf(w, "hello you")
	helper.FailOnError(err)
}

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/send-notification", sendNotificationHandler).Methods("POST")

	log.Printf("Listening on localhost:%v\n", config.Config["APP_PORT"])
	err := http.ListenAndServe(fmt.Sprintf(":%v", config.Config["APP_PORT"]), router)
	helper.FailOnError(err)
}