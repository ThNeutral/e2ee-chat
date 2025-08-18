package server

import (
	"chat/shared/logging"
	"net/http"

	"github.com/coder/websocket"
)

func (s *Server) HandleConnectChat(w http.ResponseWriter, r *http.Request) {
	logger := logging.GetLogger(r.Context())

	conn, err := websocket.Accept(w, r, nil)
	if err != nil {
		logger.Error("failed to accept connection", "err", err)
		return
	}

	err = s.hub.AddConnection(conn)
	if err != nil {
		logger.Error("failed to add connection", "err", err)
		err := conn.CloseNow()
		if err != nil {
			logger.Error("failed to close connection", "err", err)
		}
		return
	}
}
