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

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	server := RunServer()

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

func RunServer() *http.Server {

	address := flag.String("address", defaultAddress, "Listening address in format <ip>:<port>")
	flag.Parse()
	log.Printf("Starting web server on %s\n", *address)

	postRepository := storages.NewInMemoryPostRepository()
	postService := services.NewPostService(postRepository)

	router := http.NewServeMux()
	router.Handle("/ws", api.NewPostHandler(postService))
	router.Handle("/", api.NewSpaHandler("static", "index.html"))

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

	return server
}
