//? Using buffered channels
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 3) // Buffered channel, as capacity is set at declaration

	// The sender of the buffer channel will block only when there is no empty slot in
	// the channel. While the receiver will block on the channel when it's empty

	go func(c chan int) {

		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("func goroutine #%d after sending data into the channel\n", i)
		}
		close(c) // Even though it's closed, we can receive the values that are waiting in the channel

		//! You don't always need to close a channel. Only imp, when it's important to
		//! tell the receiving goroutines that all the data has been sent.
	}(c)

	fmt.Println("main goroutine sleep for 2 seconds") // To give time to other goroutine to start
	time.Sleep(time.Second * 2)

	//? To receive data from a channel, where more goroutines are data, is to use a for loop
	//? and iterate over the channel:
	for v := range c { //* Same as using: v := <- c (for each value in the channel)
		fmt.Println("main goroutine received value from channel: ", v)
	}

	fmt.Println("\n", <-c) //! Value received from closed channel = '0' value of that type

	c <- 5 //! PANICS, when sending value to closed channel
}
