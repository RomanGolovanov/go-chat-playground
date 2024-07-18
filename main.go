package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/RomanGolovanov/go-chat-playground/api"
	"github.com/RomanGolovanov/go-chat-playground/api/middleware"
	"github.com/RomanGolovanov/go-chat-playground/internal/services"
	"github.com/RomanGolovanov/go-chat-playground/internal/storages"
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

	router := http.NewServeMux()

	router.Handle("/ws", postHandler)
	router.Handle("/", spa)

	log.Printf("Starting web server on %s\n", *address)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	server := &http.Server{
		Addr: *address,
		Handler: middleware.NewMiddlewareChain(
			middleware.NewLoggingMiddleware(),
			middleware.NewCorsMiddleware(middleware.CorsOptions{
				AllowOrigin:      "*",
				AllowHeaders:     "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token",
				AllowMethods:     "POST, GET, OPTIONS, PUT, DELETE",
				AllowCredentials: "true",
			}),
		)(router),
	}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalf("Server shutdown: %v\n", err)
		}

	}()

	<-ctx.Done()

	log.Println("Shutting down gracefully, press Ctrl+C again to force")
	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")
}
