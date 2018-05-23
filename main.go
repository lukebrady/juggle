package main

import (
	"net/http"
)

func main() {
	// Read in the juggle configuration to be used by the program.
	configuration, err := ReadJuggleConfig("config/juggle.conf.json")
	if err != nil {
		panic(err)
	}
	// Check which algorithm has been used in the configuration.
	switch configuration.Algorithm {
	case "rr":

		break
	case "ll":
		break
	case "pr":
		break
	default:
		break
	}
	// Create the web multiplexer.
	server := http.NewServeMux()
	// Listen and serve the load-balancer server.
	http.ListenAndServe(configuration.Port, server)
}
