package handlers

import (
	"encoding/json"
	"net/http"
)

func httpError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write(formatJSONError(err.Error()))
}

func formatJSONError(message string) []byte {
	appError := struct {
		Message string `json:"message"`
	}{
		message,
	}
	response, err := json.Marshal(appError)
	if err != nil {
		return []byte(err.Error())
	}
	return response
}
