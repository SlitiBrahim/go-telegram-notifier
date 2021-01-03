package helper

import (
	"encoding/json"
	"log"
	"net/http"
)

func FailOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func SendApiError(w http.ResponseWriter, err error, httpStatus int) {
	if err != nil {
		httpErr := map[string]interface{}{
			"error": err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpStatus)

		json.NewEncoder(w).Encode(httpErr)
	}
}
