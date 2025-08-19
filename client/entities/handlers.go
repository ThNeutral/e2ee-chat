package entities

type OnClickHandler func()
type OnInputHandler func(Component, []rune)

type OnConnectHandler func()
type OnDisconnectHandler func()
