package main

import (
	"fmt"
	"log"
	"net/http"
)

type RequestHandler struct{}

func (rh *RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func main() {
	fmt.Println("Simple HTTP Server is running")

	handler := &RequestHandler{}

	server := http.Server{
		// Note: This is NOT secure - As we are binding to
		// all the network interfaces and any of them can be
		// part of a network that's not intended to have access
		// to the server, and this could be the public network too -
		// Internet, with public IP
		// Using port 8080 to avoid permission issues like
		// requiring sudo permission for registered ports like 80
		// This is the same as -
		// Addr: "0.0.0.0:8080",
		Addr:    ":8080",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("an error occurred while running the server: %v", err)
	}
}

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
