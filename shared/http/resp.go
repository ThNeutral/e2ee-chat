package http

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Message string `json:"message"`
	Payload any    `json:"payload"`
}

func WriteError(w http.ResponseWriter, code int, message string) {
	body := response{
		Message: message,
		Payload: nil,
	}

	bytes, _ := json.Marshal(body)

	w.WriteHeader(code)
	w.Write(bytes)
}

func WriteSuccess(w http.ResponseWriter, payload any) {
	body := response{
		Message: "",
		Payload: payload,
	}

	bytes, _ := json.Marshal(body)

	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
