package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int) // unbuffered channel

	c2 := make(chan int, 3) // buffered channel, 2nd argument = capacity of channel
	_ = c2

	go func(c chan int) {
		fmt.Println("func goroutine starts sending data into the channel")
		c <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c1)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println("main routine starts receiving data")
	d := <-c1
	fmt.Println("main goroutine received data: ", d)

	time.Sleep(time.Second)
}

// For an unbuffered channel, the sender (the func goroutine) blocks the channel, until
// the receiver(the main goroutine) receives the data from the channel
//! c <- 10  => Blocks the goroutine until main wakes up(2 secs). As it can't read from
//! the channel when it's sleeping.
