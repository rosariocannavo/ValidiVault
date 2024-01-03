package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/nats-io/nats.go"
)

var (
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)
var nc *nats.Conn
var subject = "rest_logging"
var file *os.File
var err error

func main() {
	natsURL := "nats://nats:4222"

	//open file
	file, err = os.OpenFile("message_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//connect to nats
	nc, err = nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/", serveHTML)

	fmt.Println("Server started on :8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
		data := string(msg.Data)
		fmt.Println("msg: ", data)

		if _, err := file.WriteString(string(msg.Data) + "\n"); err != nil {
			log.Println("Error writing to file:", err)
		}
		err = ws.WriteJSON(data)
		if err != nil {
			log.Printf("error: %v", err)
		}
	})

	if err != nil {
		log.Fatal(err)
		ws.Close()
	}
}

func serveHTML(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
