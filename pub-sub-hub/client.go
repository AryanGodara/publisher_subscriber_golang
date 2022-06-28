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
			log.Println("Error in receive: ", err)
			// Conn_lost <- connection //? Send current connection reference to this channel
			return
		}

		log.Printf("Received: %s\n", msg)
	}
}

func main() {
	done = make(chan interface{})    // Channel to indicate that the receiverHandler is done
	interrupt = make(chan os.Signal) // Channel to listen for interrupt signal to terminate gracefully

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
			// We received a SIGINT (Ctrl + C). Terminate gracefully...
			log.Println("Interrupt occured, closing the connection...")

			// Close our websocket connection
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
