package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// We'll need to define an Upgrader, this'll require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// define a reader which will listen for new messages being sent to our WS endpoint
func reader(conn *websocket.Conn) {
	for {
		// read in a message
		messageType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		// print out that message for clarity
		fmt.Println(string(msg))

		if err := conn.WriteMessage(messageType, msg); err != nil {
			log.Println(err)
			return
		}
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	// CheckOrigin determines whether or not an incoming request from a different domain is allowed to connect, and if it isn't they'll be hit with a CORS error
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	defer ws.Close() // good practice

	// helpful log statement to show connections
	log.Println("Client Connected via WS protocol")

	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
	}

	// listen indefinitely for new messages coming through on our WebSocket connection
	reader(ws)
}

func setupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home Page")
	})

	http.HandleFunc("/ws", wsEndpoint)
}

func main() {
	fmt.Println("Hello World!!")
	setupRoutes()

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
