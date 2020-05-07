package service

import (
	"encoding/json"
	"net/http"
)

func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func ErrorHandler(w http.ResponseWriter, err error) {
	errMsg := map[string]interface{}{"message": err.Error(), "statusCode": 400}
	RespondWithJSON(w, http.StatusBadRequest, errMsg)
}
