package main

import (
	"fmt"
	"log"
	"strconv"
)

// A type that satisfies the fmt.Stringer interface
type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s -%s", b.Title, b.Author)
}

// Another type that satisfies the fmt.String interface
type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// Declare a WriteLog() function which takes any object that satisfies the "fmt.Stringer
// interface", as a parameter
func WriteLog(s fmt.Stringer) {
	log.Print(s.String())
}

func main() {
	book := Book{"Abcd", "Aryan Godara"}
	WriteLog(book)

	// var count Count
	// count = Count(3)
	count := Count(3)
	WriteLog(count)
}
