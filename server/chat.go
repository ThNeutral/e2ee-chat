package server

import (
	"chat/shared/ctxinjector"
	"net/http"
)

func (s *Server) HandleConnectChat(w http.ResponseWriter, r *http.Request) {
	logger := ctxinjector.GetLogger(r.Context())

	err := s.hub.Accept(w, r)
	if err != nil {
		logger.Error("failed to accept connection", "err", err)
		return
	}
}
