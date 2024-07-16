package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsReader(ws *websocket.Conn) {
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

func WebSocketEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Client connected")
	wsReader(ws)
}
