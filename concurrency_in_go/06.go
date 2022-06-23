//* A channel in go provides a connection b/w goroutines, allowing them to communicate.
//* Channels are used to communicate b/w running goroutines.
//? The data we're sending or receiving from a channel, must always be of the same type
//* A channel is a two-way communicator

//! This one results in a "deadlock", since there are no goroutines, for the channel created

package main

import "fmt"

func main() {
	var ch chan int
	fmt.Println(ch)
	// 'ch' has type 'chan int'
	// we want to communicate through this channel with values of type 'int'

	ch = make(chan int) //? Initializing a channel
	fmt.Println(ch)     //* A channel is like a pointer

	// c := make(chan int)	// Shorter declaration + initialization

	// <- channel operator

	//SEND
	ch <- 10 // Send value to the channel

	//RECEIVE
	num := <-ch // Receive value from the channel
	_ = num
	fmt.Println(<-ch)

	close(ch)
}

/*
aryan@aryan-ZenBook-UX425JA-UX425JA:~/publisher_subscriber_golang/concurrency_in_go$ go run 06.go
<nil>
0xc0000220c0
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
        /home/aryan/publisher_subscriber_golang/concurrency_in_go/06.go:24 +0xbb
exit status 2
aryan@aryan-ZenBook-UX425JA-UX425JA:~/publisher_subscriber_golang/concurrency_in_go$
*/
