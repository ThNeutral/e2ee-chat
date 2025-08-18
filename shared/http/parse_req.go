package http

import (
	"encoding/json"
	"io"
	"net/http"
)

type Validatable interface {
	Validate() error
}

func ParseHTTPRequest[T Validatable](r *http.Request) (T, error) {
	var val T

	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		return val, err
	}

	err = json.Unmarshal(bytes, &val)
	if err != nil {
		return val, err
	}

	err = val.Validate()
	if err != nil {
		return val, err
	}

	return val, nil
}
