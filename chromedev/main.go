package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/hayeah/chromedev"
)

func main() {
	var port uint
	defaultPort := 9222

	// Define a flag for the port
	flag.UintVar(&port, "port", uint(defaultPort), "specify the port for remote debugging")

	// Parse the command-line flags
	flag.Parse()

	// Check if a port argument was provided
	if len(flag.Args()) > 0 {
		// Attempt to parse the provided port argument
		p, err := strconv.Atoi(flag.Args()[0])
		if err != nil {
			fmt.Printf("Invalid port argument: %v\n", err)
			os.Exit(1)
		}
		port = uint(p)
	}

	err := chromedev.Open(port)
	if err != nil {
		log.Fatalf("failed to open chrome dev tools: %v", err)
	}
}
