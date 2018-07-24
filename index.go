package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// IndexServer is the struct that keeps track of index server meta-data.
type IndexServer struct {
	protocol string
	hosts    []string
	rr       RRLinkedList
}

// NewIndexServer creates and returns an IndexServer object.
func NewIndexServer(hosts []string) *IndexServer {
	return &IndexServer{
		protocol: "http",
		hosts:    hosts,
		rr:       *NewRRLinkedList(hosts),
	}
}

// Route is the main route of the load-balancer that can load-balance traffic.
func (i *IndexServer) Route(w http.ResponseWriter, r *http.Request) {
	// If a get method has been sent to the load-balancer, send a GET
	// request to the host currently in the rr lineup.
	if r.Method == "GET" {
		data, err := http.Get(i.hosts[i.rr.current])
		if err != nil {
			panic(err)
		}
		// Now read the data given by the load-balancer and return if 200.
		html, err := ioutil.ReadAll(data.Body)
		if err != nil {
			panic(err)
		}
		// Now increment the host in the load-balancer.
		fmt.Println(i.rr.current)
		// Write the html data.
		w.Write(html)
		w.WriteHeader(http.StatusOK)
		i.rr.NextHost()
	}
}
