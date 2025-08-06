package main

import (
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("mode", "client", "specify mode of operation (server or client)")
	port := flag.Int("port", 8080, "specify port to listen for connections (only for server)")
	serverAddress := flag.String("address", "127.0.0.1:8081", "specify central server address (only for client)")

	flag.Parse()

	if *mode == "server" {
		startServer(*port)
	} else if *mode == "client" {
		startClient(*serverAddress)
	} else {
		fmt.Printf("Unknown mode of operation: %v. Allowed values are either 'server' or 'client'\n", *mode)
	}
}
