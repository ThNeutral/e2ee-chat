package server

import (
	"chat/shared"
	"encoding/json"
	"net/http"
)

type echoRequest struct {
	Value string `json:"value"`
}

func (r echoRequest) Validate() error {
	eb := shared.NewErrorBuilder().Msg("failed to validate echo request")

	if r.Value == "" {
		return eb.Causef("value cannot be empty").Err()
	}

	return nil
}

func (s *Server) handleEcho(w http.ResponseWriter, r *http.Request) {
	eb := shared.NewErrorBuilder().Msg("failed to handle echo")

	request, err := shared.ParseHTTPRequest[echoRequest](r)
	if err != nil {
		shared.WriteHTTPError(w, http.StatusBadRequest, eb.Cause(err).Err())
		return
	}

	bytes, _ := json.Marshal(request)
	w.Write(bytes)
}
