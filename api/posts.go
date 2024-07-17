package api

import (
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
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Client connected")
	reader(ws)
}

func reader(ws *websocket.Conn) {
	for {
		messageType, buff, err := ws.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(buff))
		if err := ws.WriteMessage(messageType, buff); err != nil {
			log.Println(err)
			return
		}
	}
}
