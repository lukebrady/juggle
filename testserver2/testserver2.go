package main

import (
	"net/http"
)

func index2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from testserver 2."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index2)
	http.ListenAndServe(":8081", mux)
}
