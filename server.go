package main

import (
	"chat/server"
	"chat/server/hub"
	"log"
)

func runServer() {
	hub := hub.New(hub.Config{})

	server := server.New(server.Config{
		Port: 8080,
		Hub:  hub,
	})

	err := server.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
