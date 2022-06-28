package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var done chan interface{}
var interrupt chan os.Signal

func receiveHandler(connection *websocket.Conn) {
	defer close(done)

	for {
		_, msg, err := connection.ReadMessage()
		if err != nil {
			log.Fatal("Error in receive: ", err)
			return
		}

		log.Printf("Received: %s\n", msg)
	}
}

func main() {
	done = make(chan interface{})
	interrupt = make(chan os.Signal)

	signal.Notify(interrupt, os.Interrupt) // Notify the interrupt channel for SIGINT

	socketUrl := "ws://localhost:8080" + "/socket" //* URL for websocket connection
	conn, _, err := websocket.DefaultDialer.Dial(socketUrl, nil)
	if err != nil {
		log.Fatal("Error connecting to Websocket Server: ", err)
	}

	defer conn.Close() //todo: Close the connection before exiting main goroutine

	go receiveHandler(conn)

	for {
		select {

		case <-interrupt:
			log.Println("Interrupt occured, closing the connection...")

			err := conn.WriteMessage(websocket.TextMessage, []byte("closed"))
			if err != nil {
				log.Println("Error during closing websocket: ", err)
				return
			}

			select {
			case <-done: // Close channel will give 0
				log.Println("Receiver Channel Closed! Exiting...")

			case <-time.After(time.Duration(1) * time.Second): // Did not receive anything from done channel
				log.Println("Timeout in closing receiving channel. Exiting...")
			}
			return
		}
	}
}
