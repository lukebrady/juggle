package web

import (
	"net/http"
)

// IndexServer is the struct that keeps track of index server meta-data.
type IndexServer struct {
	protocol string
	hosts []string
}

func Route(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Get()
	}
}