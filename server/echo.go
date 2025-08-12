package server

import (
	"chat/shared"
	"encoding/json"
	"fmt"
	"net/http"
)

type EchoRequest struct {
	Value string `json:"value"`
}

type EchoResponse struct {
	Value string `json:"value"`
}

func (r EchoRequest) Validate() error {
	if r.Value == "" {
		return fmt.Errorf("value cannot be empty")
	}

	return nil
}

func (s *Server) HandleEcho(w http.ResponseWriter, r *http.Request) {
	eb := shared.B().Msg("failed to handle echo")

	request, err := shared.ParseHTTPRequest[EchoRequest](r)
	if err != nil {
		shared.WriteHTTPError(w, http.StatusBadRequest, eb.Cause(err).Err())
		return
	}

	resp := EchoResponse{
		Value: request.Value,
	}

	bytes, _ := json.Marshal(resp)
	w.Write(bytes)
}
