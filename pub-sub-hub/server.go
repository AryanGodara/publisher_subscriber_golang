package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

//! HUB CODE

// Hub maintains the set of active clients and broadcasts messages to the clients
type Hub struct {
	// Register lients
	clients map[*Client]bool

	// Inbound messages from the clients
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients
	unregister chan *Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

//! HUB CODE

//! CLIENT STRUCT

type Client struct {
	hub *Hub

	conn *websocket.Conn

	send chan []byte
}

//! CLIENT STRUCT

var msg string = "Starting out message"
var flag bool = false

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
	flag = true
}

//* HTTP Portion

//* WebSocket Portion
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
} // Use default options

func socketHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// Upgrade our raw HTTP connection to a websocket based one
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation", err)
		return
	}
	defer conn.Close()

	//* Register with the hub
	client := &Client{hub: hub, conn: conn, send: make(chan []byte)}
	client.hub.register <- client

	defer func() {
		client.hub.unregister <- client
		client.conn.Close()
	}()
	//* Register with the hub

	// The event loop
	for {
		//? Delete select block to restore
		switch {
		case flag == true:
			log.Println("Sending Message: ", msg)
			err = conn.WriteMessage(2, []byte(msg)) //* 2 for binary message
			if err != nil {
				log.Println("Error during message writing", err)
				return
			}
			flag = false // Turns true when message changes
		}

	}
}

//* WebSocket Portion

func main() {
	//? For hub.go
	hub := newHub()
	go hub.run()

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
	http.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		socketHandler(hub, w, r) // To pass in the 'hub' as an argument
	})

	fmt.Println("WebSocket Server listening on port 8080...")

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err.Error())
	}
}
