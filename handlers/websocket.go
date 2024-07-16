package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

func HandleWebSocket(router *mux.Router, pathPrefix string) {
	router.PathPrefix(pathPrefix).Handler(wsHandler{})
}

type wsHandler struct{}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (h wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
