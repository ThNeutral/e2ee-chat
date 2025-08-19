package ws

import "chat/client/entities"

func (ws *Websocket) SetOnConnectHandler(handler entities.OnConnectHandler) {
	ws.onConnectHandler = handler
}

func (ws *Websocket) SetOnDisconnectHandler(handler entities.OnDisconnectHandler) {
	ws.onDisconnectHandler = handler
}
