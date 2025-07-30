package shared

import (
	"encoding/json"
	"net/http"
)

type httpError struct {
	Message string `json:"message"`
}

func WriteHTTPError(w http.ResponseWriter, code int, err error) {
	body := httpError{
		Message: err.Error(),
	}

	bytes, _ := json.Marshal(body)

	w.WriteHeader(code)
	w.Write(bytes)
}
