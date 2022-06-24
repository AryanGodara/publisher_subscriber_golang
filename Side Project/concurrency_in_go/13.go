//? CHANNEL SELECT STATEMENTS

// the select statement lets a goroutine wait on multiple communication operations.
// select blocks until one of its cases can run, then it executes that case.
// it chooses one at random, if multiple are ready.

//* TRY: for 2->3, remove/add time.sleep(), remove/add default:

package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixNano() / 1000000

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Hii!"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Bye!"
	}()

	time.Sleep(time.Second * 2)
	for i := 0; i < 2; i++ {
		select { //* Kind of like switch, but only used with channels
		case msg1 := <-c1:
			fmt.Println("Received: ", msg1)

		case msg2 := <-c2:
			fmt.Println("Received:", msg2)
		default:
			fmt.Println("No Activity")
		}
	}

	end := time.Now().UnixNano() / 1000000

	// 2000 (Both run concurrently)
	fmt.Println(end - start) // Prints the milliseconds elapsed since program started
}
