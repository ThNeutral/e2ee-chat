package main

import (
	"chat/client"
	"chat/client/repository"
	"chat/client/services/echo"
	"chat/raylib"
	"chat/raylib/entities"
	"chat/raylib/runner"
	"log"
	"net"
	"net/http"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
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

	cl, err := client.New(client.Config{
		ServerAddr: addr,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},

		DefaultTimeout: 10 * time.Second,

		Echo: echo,
	})
	if err != nil {
		log.Fatalln(err)
	}

	runner := runner.New(runner.Config{
		WindowConfig: entities.WindowConfig{
			Width:           800,
			Height:          600,
			Title:           "TEST",
			BackgroundColor: rl.DarkGreen,
		},
	})
	runner.SetupInitialLayout()

	raylib := raylib.New(raylib.Config{
		Runner: runner,

		Echo: cl,
	})

	log.Println("Started client")
	raylib.Run()
}
