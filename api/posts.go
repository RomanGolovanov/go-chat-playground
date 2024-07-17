package api

import (
	"context"
	"log"
	"net/http"

	"github.com/RomanGolovanov/go-chat-playground/internal/services"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type PostHandler struct {
	service *services.PostService
}

func NewPostHandler(service *services.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func HandlePostsWebSocket(router *mux.Router, pathPrefix string, handler *PostHandler) {
	router.PathPrefix(pathPrefix).Handler(handler)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h PostHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Client connected")

	posts, err := h.service.GetPosts(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Replay posts")
	for _, post := range posts {
		buff := ([]byte)(post.Text)
		ws.WriteMessage(1, buff)
	}

	log.Println("Run message relay")
	reader(ctx, h, ws)
}

func reader(ctx context.Context, h PostHandler, ws *websocket.Conn) {
	for {
		messageType, buff, err := ws.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		text := string(buff)
		log.Println(text)

		if len(text) != 0 {
			h.service.AddPost(ctx, services.CreatePostRequest{From: "unknown", Text: text})
			if err := ws.WriteMessage(messageType, buff); err != nil {
				log.Println(err)
				return
			}
		}
	}
}
