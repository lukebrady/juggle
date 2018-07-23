package main

import (
	"fmt"
	"net/http"
)

func indexRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.Write([]byte("Juggle v1: The programmable load-balancer."))
	}
}

func main() {
	// Read in the juggle configuration to be used by the program.
	configuration, err := ReadJuggleConfig("./config/juggle.conf.json")
	if err != nil {
		panic(err)
	}
	// Check which algorithm has been used in the configuration.
	/*
		switch configuration.Algorithm {
		case "rr":
			break
		case "ll":
			break
		case "pr":
			break
		default:
			break
		}  */
	// Create the web multiplexer.
	server := http.NewServeMux()
	// Handle the index route.
	server.HandleFunc("/", indexRoute)
	fmt.Println(configuration.Port)
	// Listen and serve the load-balancer server.
	http.ListenAndServe(":9080", server)
}
