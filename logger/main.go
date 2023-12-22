package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to the NATS server, use docker network DNS name
	natsURL := "nats://nats:4222"
	nc, err := nats.Connect(natsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	file, err := os.OpenFile("message_log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Subject to subscribe to
	subject := "rest_logging"

	// Subscribe to the subject and read the messages when up
	_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
		fmt.Printf("%s", string(msg.Data))
		if _, err := file.WriteString(string(msg.Data) + "\n"); err != nil {
			log.Println("Error writing to file:", err)
		}
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Subscriber connected and listening...")

	// Keep the program running
	select {}
}
