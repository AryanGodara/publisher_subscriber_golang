package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello")
}

func world(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "world")
}

func main() {

	serverMuxA := http.NewServeMux() // ServeMuxA
	serverMuxA.HandleFunc("/", hello)

	serverMuxB := http.NewServeMux() // ServeMuxB
	serverMuxB.HandleFunc("/", world)

	http.HandleFunc("/", helloHandler) // Default ServeMux

	go func() {
		fmt.Println("Listening on 8080")
		err := http.ListenAndServe(":8080", serverMuxA)
		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}()

	go func() {
		fmt.Println("Listening on 8088")
		err := http.ListenAndServe(":8088", serverMuxB)
		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}()

	go func() {
		fmt.Println("Listening on 8888")
		err := http.ListenAndServe(":8888", nil)
		if err != nil {
			panic("ListenAndServe: " + err.Error())
		}
	}()

	fmt.Println("All 3 servers have started in different goroutines...")

	select {} // To block the main goroutine while the other two are active
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
