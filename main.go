package main

import (
	"chat/client"
	"chat/shared"
	"flag"
	"log"
	"net"
)

func main() {
	eb := shared.NewErrorBuilder().Msg("failed to initialize app")

	fePort := flag.Int("port", 8080, "specify port to listen for front end connections")
	serverAddress := flag.String("address", "127.0.0.1:8081", "specify central server address")

	flag.Parse()

	log.Printf("Provided FE listen port is %v\n", *fePort)
	log.Printf("Provided cental server address is %v\n", *serverAddress)

	addr, err := net.ResolveTCPAddr("tcp", *serverAddress)
	if err != nil {
		log.Fatalln(eb.Cause(err).Err())
	}

	cl, err := client.New(addr, *fePort)
	if err != nil {
		log.Fatalln(eb.Cause(err).Err())
	}

	log.Println("Started client")
	err = cl.Run()
	if err != nil {
		log.Fatalln(eb.Cause(err).Err())
	}
}
