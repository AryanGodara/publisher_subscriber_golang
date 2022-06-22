package main

import (
	"fmt"
	"net/http"
)

func BasicHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my website")
}

func main() {
	http.HandleFunc("/", BasicHandler)
	http.ListenAndServe(":8080", nil)
}
