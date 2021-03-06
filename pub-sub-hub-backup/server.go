package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Conn_lost = make(chan *websocket.Conn)

//! HUB CODE

// Hub maintains the set of active clients and broadcasts messages to the clients
type Hub struct {
	// Register lients
	clients map[*Client]bool

	// Inbound messages from the http server
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients
	unregister chan *Client
}

func newHub() *Hub { //? This one can be in the main function itself
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() { //* Continuosly listening for registeration from clients, and messages from http server
	for {
		select {
		case con := <-Conn_lost:
			//? We've found a closed client connection. Match this with all clients in map
			// for client := range h.clients {
			// 	if h.clients[client] == true && client.conn == con {
			// 		delete(h.clients, client)
			// 		fmt.Println("Client Unregistered Forcefully")
			// 	}
			// }
			h.remove_closed(con)
		case client := <-h.register:
			h.clients[client] = true
			fmt.Println("Client Registered")
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				fmt.Println("Client Unregistered")
				delete(h.clients, client)
			}
		case message := <-h.broadcast:
			fmt.Println("Message received to hub: ", string(message))
			//! Before sending, check status of all clients
			select {
			case con := <-Conn_lost:
				h.remove_closed(con)
			default:
				for client := range h.clients {
					err := client.conn.WriteMessage(2, []byte(msg)) //* 2 for binary message
					if err != nil {
						log.Println("Error during message writing", err)
						return
					}
					// select {
					// case client.send <- message:
					// default:
					// 	close(client.send)
					// 	delete(h.clients, client)
					// }
				}
			}

			// for client := range h.clients {
			// 	err := client.conn.WriteMessage(2, []byte(msg)) //* 2 for binary message
			// 	if err != nil {
			// 		log.Println("Error during message writing", err)
			// 		return
			// 	}
			// 	// select {
			// 	// case client.send <- message:
			// 	// default:
			// 	// 	close(client.send)
			// 	// 	delete(h.clients, client)
			// 	// }
			// }
		}
	}
}

//? Interrupt client unregister function
func (h *Hub) remove_closed(con *websocket.Conn) {
	for client := range h.clients {
		if h.clients[client] == true && client.conn == con {
			delete(h.clients, client)
			fmt.Println("Client Unregistered Forcefully")
		}
	}
}

//! HUB CODE

//! CLIENT STRUCT

type Client struct {
	hub *Hub

	conn *websocket.Conn
}

//! CLIENT STRUCT

var msg string = "Starting out message"

//* HTTP Portion
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Index Page")
}

func httppostHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Use POST method here", http.StatusMethodNotAllowed)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	msg = string(b) //* Uncomment these two lines to revert back
	// flag = true

	//* Comment below code to revert back
	if len(hub.clients) != 0 {
		hub.broadcast <- []byte(msg) //* BROADCASTING THE MESSAGE
		fmt.Fprintln(w, "This is the message: ", string(b))
	} else {
		fmt.Println("Message received, but no place to send")
		fmt.Fprintln(w, "Message received, but no place to send")
	}
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

	//* Register and (defer Unregister) with the hub
	// client := &Client{hub: hub, conn: conn, send: make(chan []byte)}
	client := &Client{hub: hub, conn: conn}
	client.hub.register <- client

	defer func() {
		client.hub.unregister <- client
		client.conn.Close()
	}()
	//* Register and (defer Unregister) with the hub

	// The event loop
	for {
		// //? If this receives it first
		// case con := <-Conn_lost:
		// 	if con == client.conn { // This is the connection, to close properly
		// 		return // This will call the defer func above.
		// 	}
		//? Receive a message "closed" from client, then push conn to conn_lost channel
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during reading message from client: ", err.Error())
			break //* Will automatically call the defer func now
		}
		if string(msg) == "closed" {
			Conn_lost <- conn
			break
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
	httpMux.HandleFunc("/postmessage", func(w http.ResponseWriter, r *http.Request) {
		httppostHandler(hub, w, r)
	})

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
