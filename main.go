package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/RomanGolovanov/go-chat-playground/api"
	"github.com/RomanGolovanov/go-chat-playground/internal/services"
	"github.com/RomanGolovanov/go-chat-playground/internal/storages"
	"github.com/gorilla/mux"
)

const (
	defaultAddress = "0.0.0.0:8080"
)

func main() {
	address := flag.String("address", defaultAddress, "Listening address in format <ip>:<port>")
	flag.Parse()

	spa := api.NewSpaHandler("static", "index.html")

	postRepository := storages.NewInMemoryPostRepository()
	postService := services.NewPostService(postRepository)
	postHandler := api.NewPostHandler(postService)

	router := mux.NewRouter()

	api.HandlePostsWebSocket(router, "/ws", postHandler)
	api.HandleSpa(router, "/", spa)

	log.Printf("Starting web server on %s\n", *address)
	err := http.ListenAndServe(*address, router)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
