package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", handleHelloRequest)
	http.ListenAndServe(":8080", mux)
}

func handleHelloRequest(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(200)
	_, _ = w.Write([]byte("Hello, world!"))
}
