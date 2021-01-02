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

	// TODO: send HTTP request to TG api

	res := map[string]interface{}{
		"response": "Notification has been sent.",
	}

	err = ReturnResponse(w, res, http.StatusOK)
	helper.SendApiError(w, err, http.StatusInternalServerError)
}

func Start() {
	router := mux.NewRouter()
	router.HandleFunc("/send-notification", sendNotificationHandler).Methods("POST")

	log.Printf("Listening on localhost:%v\n", config.Config["APP_PORT"])
	err := http.ListenAndServe(fmt.Sprintf(":%v", config.Config["APP_PORT"]), router)
	helper.FailOnError(err)
}

func ReturnResponse(w http.ResponseWriter, res map[string]interface{}, httpStatus int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(httpStatus)

	err := json.NewEncoder(w).Encode(res)

	return err
}
