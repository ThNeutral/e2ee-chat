package main

import (
	"chat/server"
	"chat/server/services"
	"chat/shared"
	"chat/shared/endpoints"
	"fmt"
	"log"
	"net/http"
)

func startServer(serverPort int) {
	eb := shared.NewErrorBuilder().Msg("failed to start client")

	log.Printf("Provided server listen port is %v\n", serverPort)

	hub := services.NewHub(services.HubConfig{})

	serv, err := server.New(server.ServerConfig{
		Hub: hub,
	})
	if err != nil {
		log.Fatalln(eb.Cause(err).Err())
	}

	runner := shared.NewRunner()
	runner.Post(endpoints.Echo, serv.HandleEcho)

	log.Println("Started server")
	err = http.ListenAndServe(fmt.Sprintf(":%v", serverPort), runner)
	if err != nil {
		log.Fatalln(eb.Cause(err).Err())
	}
}
