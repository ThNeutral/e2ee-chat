package main

import (
	"chat/client"
	"chat/client/entities"
	"chat/client/repository"
	"chat/client/services/echo"
	"chat/client/services/raylib"
	"chat/shared"
	"log"
	"net"
	"net/http"
	"time"
)

func startClient(serverAddress string) {
	log.Printf("Provided cental server address is %v\n", serverAddress)

	addr, err := net.ResolveTCPAddr("tcp", serverAddress)
	if err != nil {
		log.Fatalln(err)
	}

	repo, err := repository.New(repository.RepositoryConfig{
		ServerAddr: addr,
		HTTPClient: &http.Client{},
	})
	if err != nil {
		log.Fatalln(err)
	}

	echo, err := echo.New(echo.Config{
		EchoRepository: repo,
	})
	if err != nil {
		log.Fatalln(err)
	}

	raylib, err := raylib.New(raylib.Config{
		WindowConfig: entities.WindowConfig{
			Width:  800,
			Height: 600,
			Title:  "TEST",
		},
	})

	cl, err := client.New(client.Config{
		ServerAddr: addr,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},

		DefaultTimeout: 10 * time.Second,

		GUI:  raylib,
		Echo: echo,
	})
	if err != nil {
		log.Fatalln(err)
	}

	runner := shared.NewRunner()
	runner.Post("/echo", cl.HandleEcho)

	log.Println("Started client")
	err = cl.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
