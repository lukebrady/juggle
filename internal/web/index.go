package web

import (
	"net/http"
)

type Index struct {
	protocol string
	hosts []string
}

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Get()
	}
}