package ws

func (ws *Websocket) finalize() {
	ws.conn.CloseNow()
	ws.conn = nil
}
