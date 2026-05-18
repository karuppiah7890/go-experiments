package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Simple HTTP Server is running")

	mux := http.NewServeMux()

	mux.HandleFunc("/livez", livenessHandler)

	server := http.Server{
		// For security purposes, only allow listening on
		// localhost (127.0.0.1) loopback network interface
		// and hence only allowing local clients to talk to the
		// server
		// Using port 8080 to avoid permission issues like
		// requiring sudo permission for registered ports like 80
		Addr:    "127.0.0.1:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatalf("an error occurred while running the server: %v", err)
	}
}

func livenessHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
