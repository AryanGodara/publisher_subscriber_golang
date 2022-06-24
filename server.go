package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var msg string = "Starting out message"

//* HTTP Portion
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Page")
}

func httppostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Use POST method here", http.StatusMethodNotAllowed)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(b))
	fmt.Fprintln(w, "This is the message: ", string(b))
	msg = string(b)
}

//* HTTP Portion

//* WebSocket Portion
var upgrader = websocket.Upgrader{} // Use default options

func socketHandler(w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation", err)
		return
	}
	defer conn.Close()

	// The event loop
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during message reading", err)
			break
		}

		log.Printf("Received: %s", message)

		// err = conn.WriteMessage(messageType, message)
		err = conn.WriteMessage(messageType, []byte(msg))
		if err != nil {
			log.Println("Error during message writing", err)
			break
		}
	}
}

//* WebSocket Portion

func main() {

	//todo: HTTP server runs on a new servemux
	httpMux := http.NewServeMux()
	httpMux.HandleFunc("/", home)
	httpMux.HandleFunc("/postmessage", httppostHandler)

	go func() {
		fmt.Println("HTTP Server listening on 4000")
		err := http.ListenAndServe(":4000", httpMux)
		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}()

	//todo: Websocket server runs on the default servemux
	http.HandleFunc("/socket", socketHandler)

	fmt.Println("WebSocket Server listening on port 8080...")

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err.Error())
	}
}
