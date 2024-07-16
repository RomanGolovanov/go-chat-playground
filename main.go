package main

import (
	"log"
	"net/http"

	"github.com/RomanGolovanov/go-chat-playground/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /", handlers.RootEndpoint)
	mux.HandleFunc("GET /ws", handlers.WebSocketEndpoint)
	address := "0.0.0.0:8080"
	log.Printf("Starting web server on %s\n", address)
	err := http.ListenAndServe(address, mux)
	if err != nil {
		log.Println(err.Error())
	}
}
