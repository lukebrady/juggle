package main

import (
	"net/http"
)

func main() {
	// Read in the juggle configuration to be used by the program.
	configuration, err := ReadJuggleConfig("./config/juggle.conf.json")
	if err != nil {
		panic(err)
	}
	// Create the IndexServer.
	indexServer := NewIndexServer(configuration.Hosts)
	// Create the web multiplexer.
	server := http.NewServeMux()
	// Handle the index route.
	server.HandleFunc("/", indexServer.Route)
	// Listen and serve the load-balancer server.
	http.ListenAndServe(configuration.Port, server)
}
