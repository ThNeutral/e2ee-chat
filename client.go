package main

import (
	"chat/client"
	"chat/client/entities"
	"chat/client/services"
	"chat/shared"
	"log"
	"net"
	"net/http"
)

func startClient(serverAddress string) {
	eb := shared.NewErrorBuilder().Msg("failed to start client")

	log.Printf("Provided cental server address is %v\n", serverAddress)

	addr, err := net.ResolveTCPAddr("tcp", serverAddress)
	if err != nil {
		log.Fatalln(eb.Cause(err).Err())
	}

	raylib, err := services.NewRaylib(services.RaylibConfig{
		WindowConfig: entities.WindowConfig{
			Width:  800,
			Height: 600,
			Title:  "TEST",
		},
	})

	cl, err := client.New(client.Config{
		ServerAddr: addr,
		HTTPClient: &http.Client{},
		GUI:        raylib,
	})
	if err != nil {
		log.Fatalln(eb.Cause(err).Err())
	}

	runner := shared.NewRunner()
	runner.Post("/echo", cl.HandleEcho)

	log.Println("Started client")
	err = cl.Run()
	if err != nil {
		log.Fatalln(eb.Cause(err).Err())
	}
}
