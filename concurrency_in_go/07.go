//? Basic example of a channel communicating b/w 2 goroutines: 1. main(), 2. f1()
package main

import "fmt"

// n = value to send to channel, ch = channel (bidirectional)
func f1(n int, ch chan int) {
	ch <- n // send n into the channel
}

func main() {
	defer fmt.Println("Exiting main...")

	c := make(chan int) // Bidirectional channel (send and receive)

	// only for receiving
	c1 := make(<-chan string) // Unidirectional channel (can only receive, not send value)

	// only for sending
	c2 := make(chan<- string)

	fmt.Printf("%T, %T, %T\n", c, c1, c2)

	go f1(10, c)

	n := <-c // Receive the value from c, that c received in f1()

	fmt.Println(n)
}
