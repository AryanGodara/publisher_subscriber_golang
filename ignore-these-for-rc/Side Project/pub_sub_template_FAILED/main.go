package main

import (
	"fmt"
	"log"
	"net/http"
)

type Client struct {
	name string
}

var ClientsSlice []Client

type Message struct {
	// Client
	topic string
	value string
}

var MessagesSlice []Message

// func baseHandler(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.Error(w, "404 not found", http.StatusNotFound)
// 		return
// 	}
// 	if r.Method != "GET" {
// 		http.Error(w, "method is not supported", http.StatusNotFound)
// 		return
// 	}
// 	fmt.Fprintln(w, "Go to \"/createClient\" to add a new client")
// 	fmt.Fprintln(w, "Go to \"/createMessage\" to add a new client")
// }

func createClientFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/createClientForm" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	fmt.Fprintf(w, "Name = %s\n", name)

	temp_struct := Client{name: string(name)}

	ClientsSlice = append(ClientsSlice, temp_struct)

	fmt.Fprintln(w, MessagesSlice)
}

func createMessageFormHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/createMessageForm" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")
	topic := r.FormValue("topic")
	value := r.FormValue("bodyvalue")
	fmt.Fprintf(w, "Name = %s\n", topic)
	fmt.Fprintf(w, "Address = %s\n", value)

	MessagesSlice = append(MessagesSlice, Message{string(topic), string(value)})

	fmt.Fprintln(w, MessagesSlice)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/form" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	fmt.Fprintf(w, "POST request successful\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	fmt.Println(fileServer)

	http.Handle("/", fileServer)
	// http.HandleFunc("/", baseHandler)

	http.HandleFunc("/createClientForm", createClientFormHandler)
	http.HandleFunc("/createMessageForm", createMessageFormHandler)

	http.HandleFunc("/form", formHandler)

	fmt.Printf("Server starting at port 8080\n")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
