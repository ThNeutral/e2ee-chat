package main

import (
	"chat/client"
	"chat/client/raylib"
	"chat/client/ws"
	"chat/shared/rlutils"
	"log"
	"net/url"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func runClient() {
	u, err := url.Parse("http://localhost:8080/chat")
	if err != nil {
		log.Fatalln(err)
	}

	ws := ws.New(ws.Config{
		WSEndpoint: u,
	})

	rl := raylib.New(raylib.Config{
		Size: rlutils.Vector2{
			X: 800,
			Y: 600,
		},
		BackgroundColor: rl.LightGray,
		WindowName:      "Window",
		TargetFramerate: 60,
	})

	cl := client.New(client.Config{
		GUI:       rl,
		Websocket: ws,
	})

	cl.Init()
	err = cl.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
