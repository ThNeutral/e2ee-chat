package main

import (
	"chat/client"
	"chat/client/raylib"
	"chat/shared/rlutils"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func runClient() {
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
		GUI: rl,
	})

	err := cl.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
