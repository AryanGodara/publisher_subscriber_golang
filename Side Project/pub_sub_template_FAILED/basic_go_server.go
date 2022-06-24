package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connected to base handler")
	fmt.Fprintln(w, "You're here again, move forward from this plz")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 3)
	fmt.Fprintln(w, sum)
	fmt.Println("Connected to about handler")
	fmt.Fprintln(w, "You're here again at ABOUT, move forward from this plz")

}

func AddValues(x, y int) int {
	return x + y
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
