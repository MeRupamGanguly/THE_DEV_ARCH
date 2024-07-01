package main

import (
	domain "THE_DEV_ARCH/Domain"
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"time"

	"github.com/gorilla/websocket"
)

/*
Add Read/Write Buffer size by Upgrader-Upgrade upgrades the HTTP server connection to the WebSocket protocol.
Create Connection by Call Upgrader with request response headers.
ReadMessage() for waiting client to Connect.
infinite loop for sending Tick by WriteMessage.
WriteMessage need ByteArray. JsonEncoder converts json to ByteArray.
*/
func tickProducer(res http.ResponseWriter, req *http.Request) {
	ws := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	connection, err := ws.Upgrade(res, req, nil)
	if err != nil {
		log.Print(err)
		return
	}
	//waiting for Client to Connect and Sending first Request
	msgTyp, msg, err := connection.ReadMessage()
	if err != nil {
		log.Print(err)
		return
	}
	log.Print("Msg Rcved ", string(msg))
	// Sending Feed to Client after first Request
	price := 56700.0
	for {
		price = price + 0.5
		tick := domain.Ticker{
			Ltp:       float64(price),
			Symbol:    "BTCUSDT:BYBNC",
			Type:      "CRYP",
			Timestamp: time.Now(),
		}
		payload := new(bytes.Buffer)
		json.NewEncoder(payload).Encode(tick)
		err = connection.WriteMessage(msgTyp, payload.Bytes())
		if err != nil {
			log.Print(err)
			return
		}
		time.Sleep(time.Second)
	}
}
func main() {
	http.HandleFunc("/ws/BTCUSDT", tickProducer)
	http.ListenAndServe(":3000", nil)
}
