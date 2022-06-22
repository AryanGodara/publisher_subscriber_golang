package main

import (
	"fmt"
	"net/http"
)

func httpHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connected to handler")
	fmt.Fprintln(w, "You're here again, move forward from this plz")
}

func main() {
	http.HandleFunc("/", httpHandler)
	// funfun := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Println("Trying handleRFunc")
	// })

	// http.HandleFunc("/", funfun)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
