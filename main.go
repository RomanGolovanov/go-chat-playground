package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/RomanGolovanov/go-chat-playground/handlers"
	"github.com/gorilla/mux"
)

const (
	defaultAddress = "0.0.0.0:8080"
)

func main() {
	address := flag.String("address", defaultAddress, "Listening address in format <ip>:<port>")
	flag.Parse()

	router := mux.NewRouter()

	handlers.HandleWebSocket(router, "/ws")
	handlers.HandleSpa(router, "/", "static", "index.html")

	log.Printf("Starting web server on %s\n", *address)
	err := http.ListenAndServe(*address, router)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
