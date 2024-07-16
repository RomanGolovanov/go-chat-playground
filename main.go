package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func setupServer() {
	http.HandleFunc("/", rootEndpoint)
	http.HandleFunc("/ws", wsEndpoint)
}

func rootEndpoint(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

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
		log.Println(messageType)

		if err := ws.WriteMessage(messageType, buff); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client connected")

	wsReader(ws)
}

func main() {
	setupServer()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
