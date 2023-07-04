package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func abc(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 page not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		fmt.Fprintln(w, "Welcome to HomePage")
	case "POST":
		b, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(b))
		fmt.Fprintln(w, "This is the message: ", string(b))
	default:
		http.Error(w, "Only get and post allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", abc)

	fmt.Printf("Starting server for testing")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
