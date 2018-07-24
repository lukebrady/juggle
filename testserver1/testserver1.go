package main

import (
	"net/http"
)

func index1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from testserver 1."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index1)
	http.ListenAndServe(":8080", mux)
}
