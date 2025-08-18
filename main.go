package main

import (
	"flag"
	"fmt"
)

func main() {
	mode := flag.String("mode", "client", "provide app running mode (\"client\" or \"server\")")

	flag.Parse()

	if *mode == "client" {
		runClient()
	} else if *mode == "server" {

	} else {
		fmt.Printf("unexpected mode: %s\nUsage:\n", *mode)
		flag.PrintDefaults()
	}
}
