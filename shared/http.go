package shared

import (
	"encoding/json"
	"io"
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

type validatable interface {
	Validate() error
}

func ParseHTTPRequest[T validatable](r *http.Request) (T, error) {
	eb := NewErrorBuilder().Msg("failed to parse http request")
	var val T

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return val, eb.Cause(err).Err()
	}

	err = json.Unmarshal(bytes, &val)
	if err != nil {
		return val, eb.Cause(err).Err()
	}

	err = val.Validate()
	if err != nil {
		return val, eb.Cause(err).Err()
	}

	return val, nil
}
