package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Message struct {
	Body string
}

func messageReader(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Use POST method here", http.StatusMethodNotAllowed)
		return
	}
	var m Message

	// Try to decode the request body into the struct. If there is an error, resond to the
	// client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the Person struct
	fmt.Fprintf(w, "New Message: %+v", m)
}

func nostructrader(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Use POST method here", http.StatusMethodNotAllowed)
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
	fmt.Fprintln(w, "This is the message: ", string(b))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/newmessage", messageReader)
	mux.HandleFunc("/mes", nostructrader)

	err := http.ListenAndServe(":4000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
