package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func basicHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()                   // parse arguments, you have to call this by yourself
	fmt.Println("r.Form: ", r.Form) // print form info in server side
	fmt.Println("path(r.URL.Path): ", r.URL.Path)
	fmt.Println("scheme(r.URL.Scheme): ", r.URL.Scheme)
	fmt.Println("(r.Form[\"url_long\"]", r.Form["url_long"])

	for k, v := range r.Form {
		fmt.Println("Key:", k)
		fmt.Println("Val:", strings.Join(v, ""))
	}

	fmt.Fprintf(w, "Hello Aryan!") // sends data to client side
}

func main() {
	http.HandleFunc("/", basicHandler)       // set router
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
