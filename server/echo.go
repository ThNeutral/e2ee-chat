package server

import (
	"chat/shared"
	"encoding/json"
	"log"
	"net/http"
)

type handleEchoRequest struct {
	Value string `json:"value"`
}

func (r handleEchoRequest) Validate() error {
	eb := shared.NewErrorBuilder().Msg("failed to validate echo request")

	log.Println(r)

	if r.Value == "" {
		return eb.Causef("value cannot be empty").Err()
	}

	return nil
}

func (s *Server) handleEcho(w http.ResponseWriter, r *http.Request) {
	eb := shared.NewErrorBuilder().Msg("failed to handle echo")

	request, err := shared.ParseHTTPRequest[handleEchoRequest](r)
	if err != nil {
		shared.WriteHTTPError(w, http.StatusBadRequest, eb.Cause(err).Err())
		return
	}

	bytes, _ := json.Marshal(request)
	w.Write(bytes)
}
