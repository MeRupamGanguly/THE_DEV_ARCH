package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

func main() {
	u := url.URL{Scheme: "ws", Host: "localhost:3500", Path: "ticker/BTCUSDT"}
	log.Printf("connecting to %s", u.String())

	connection, _, err := websocket.DefaultDialer.Dial("ws://localhost:3000/ws/BTCUSDT", nil)
	if err != nil {
		return
	}
	payload, err := json.Marshal("hii")
	if err != nil {
		return
	}
	connection.WriteMessage(websocket.TextMessage, payload)
	for {
		_, resMessage, err := connection.ReadMessage()
		if err != nil {
			return
		}
		fmt.Print(string(resMessage))
	}
}
