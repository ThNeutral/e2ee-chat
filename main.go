package main

import (
	"chat/client"
	"chat/server"
	"chat/shared"
	"flag"
	"fmt"
	"log"
	"net"
)

func main() {
	mode := flag.String("mode", "client", "specify mode of operation (server or client)")
	port := flag.Int("port", 8080, "specify port to listen for connections")
	serverAddress := flag.String("address", "127.0.0.1:8081", "specify central server address (only for client)")

	flag.Parse()

	if *mode == "server" {
		startServer(*port)
	} else if *mode == "client" {
		startClient(*serverAddress, *port)
	} else {
		fmt.Printf("Unknown mode of operation: %v. Allowed values are either 'server' or 'client'\n", *mode)
	}
}

func startServer(serverPort int) {
	eb := shared.NewErrorBuilder().Msg("failed to start client")

	log.Printf("Provided server listen port is %v\n", serverPort)

	serv := server.New(serverPort)

	log.Println("Started server")
	err := serv.Run()
	if err != nil {
		log.Fatalln(eb.Cause(err).Err())
	}
}

func startClient(serverAddress string, clientPort int) {
	eb := shared.NewErrorBuilder().Msg("failed to start client")

	log.Printf("Provided client listen port is %v\n", clientPort)
	log.Printf("Provided cental server address is %v\n", serverAddress)

	addr, err := net.ResolveTCPAddr("tcp", serverAddress)
	if err != nil {
		log.Fatalln(eb.Cause(err).Err())
	}

	cl, err := client.New(addr, clientPort)
	if err != nil {
		log.Fatalln(eb.Cause(err).Err())
	}

	log.Println("Started client")
	err = cl.Run()
	if err != nil {
		log.Fatalln(eb.Cause(err).Err())
	}
}
