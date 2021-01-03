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

func authenticatedReq(r *http.Request) bool {
	return r.Header.Get("token") == config.Config["TOKEN"]
}

func sendNotificationHandler(w http.ResponseWriter, r *http.Request) {
	if authenticatedReq(r) == false {
		helper.SendApiError(w, errors.New("invalid token"), http.StatusForbidden)
		return
	}

	var notification Notification
	err := json.NewDecoder(r.Body).Decode(&notification)
	if err != nil {
		helper.SendApiError(w, errors.New("invalid request body: cannot parse body to Notification object"), http.StatusBadRequest)
		return
	}

	if notification.Message == "" {
		helper.SendApiError(w, errors.New("empty message passed"), http.StatusBadRequest)
		return
	}

	msg := Message{
		ChatID: config.Config["TG_CHAT_ID"].(string),
		Text:   notification.Message,
	}

	telegramResponse, err := sendMessage(msg)
	helper.SendApiError(w, err, http.StatusInternalServerError)

	if telegramResponse.StatusCode == http.StatusOK {
		res := map[string]interface{}{
			"message": "Notification has been sent.",
		}

		err = ReturnResponse(w, res, http.StatusOK)
		helper.SendApiError(w, err, http.StatusInternalServerError)
	} else {
		res := map[string]interface{}{
			"message": "Notification cannot be sent.",
			"error":   err.Error(),
		}

		err = ReturnResponse(w, res, http.StatusBadRequest)
		helper.SendApiError(w, err, http.StatusInternalServerError)
	}
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
