package main

import (
	"fmt"
	"log"
	"net/http"
)

// By wrapping the handler func inside another function. We give the handler func access
// to the variable, we passed as the parameters for outer func. This way, I can maybe also
// send custom messages -> A shortcut of using post requests :) hopefully
func postHandler(input_str string) http.Handler {
	// fn := func(w http.ResponseWriter, r *http.Request) {
	// 	tm := time.Now().Format(format)
	// 	fmt.Fprintln(w, "The time is: ", tm)
	// }

	// return http.HandlerFunc(fn) // Converts function into a handler function

	// OR EVEN :-
	/*
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){..})
	*/
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, r.Method)
		fmt.Fprintln(w, input_str)
	})
}

func main() {
	// Use the http.NewServeMux() function to create an empty servemux
	mux := http.NewServeMux()

	// Use the http.RedirectHandler() function to create an empty servemux
	rh := http.RedirectHandler("http://example.org", http.StatusTemporaryRedirect)

	// Next, we use the mux.Handle() function to register this with our new servemux, so
	// it acts as the handler for all incoming requests with the URL path /foo
	mux.Handle("/foo", rh)

	th := postHandler("Anything I want") // Passed in a string argument, inside the function
	// This will now return -> th is of type http.handler
	//* So, we can pass it in mux.handle()

	mux.Handle("/echo", th)

	log.Println("Listening on port 3000...")

	// Then we create a new server, and start listening for incoming requests with the
	// http.ListenAndServe() function, passing in our servemux for it to match requests
	// against as the second parameter
	http.ListenAndServe(":3000", mux)
}
