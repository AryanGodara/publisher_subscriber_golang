package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var Conn_lost = make(chan *websocket.Conn)
var msg string = "Starting out message"

type Hub struct {
	clients    map[*Client]bool
	broadcast  chan []byte
	register   chan *Client
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
		case con := <-Conn_lost:
			h.remove_closed(con)

		case client := <-h.register:
			h.clients[client] = true
			log.Println("Client Registered")

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				log.Println("Client Unregistered")
				delete(h.clients, client)
			}

		case message := <-h.broadcast:
			log.Println("Message received to hub: ", string(message))
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
				}
			}
		}
	}
}

//? Interrupt client unregister function
func (h *Hub) remove_closed(con *websocket.Conn) {
	for client := range h.clients {
		if h.clients[client] == true && client.conn == con {
			client.conn.Close()
			delete(h.clients, client)
			log.Println("Client Unregistered Forcefully")
		}
	}
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
}

func httpHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Use POST method here", http.StatusMethodNotAllowed)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	msg = string(b)

	if len(hub.clients) != 0 {
		hub.broadcast <- []byte(msg)
		fmt.Fprintln(w, "This is the message: ", string(b))
	} else {
		log.Println("Message received, but no place to send")
		fmt.Fprintln(w, "Message received, but no place to send")
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("Error during connection upgradation", err)
		return
	}

	defer conn.Close()

	client := &Client{hub: hub, conn: conn}
	client.hub.register <- client

	defer func() {
		client.hub.unregister <- client
		client.conn.Close()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error during reading message from client: ", err.Error())
			break
		}

		if string(msg) == "closed" {
			Conn_lost <- conn
			break
		}
	}
}

func main() {
	hub := newHub()
	go hub.run()

	httpMux := http.NewServeMux()

	httpMux.HandleFunc("/message", func(w http.ResponseWriter, r *http.Request) {
		httpHandler(hub, w, r)
	})

	go func() {
		log.Println("HTTP Server listening on 4000")
		err := http.ListenAndServe(":4000", httpMux)
		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}()

	http.HandleFunc("/socket", func(w http.ResponseWriter, r *http.Request) {
		wsHandler(hub, w, r) // To pass in the 'hub' as an argument
	})

	log.Println("WebSocket Server listening on port 8080...")

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe Error: ", err.Error())
	}
}
