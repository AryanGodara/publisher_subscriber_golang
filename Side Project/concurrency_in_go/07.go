package main

import "fmt"

func factorial(n int, c chan int) {
	f := 1

	for i := 2; i <= n; i++ {
		f *= i
	}

	c <- f // Send the value to the channel
}

func main() {
	ch := make(chan int)
	defer close(ch)

	go factorial(5, ch)

	// After the factorial goroutine is started, main should wait for a message to come
	// to the channel. This is called a blocking call.
	// The main goroutine is put to sleep and waits for f1 goroutine to send the message
	// to the channel
	f := <-ch
	fmt.Println(f)

	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		f := <-ch
		fmt.Println(i, " factorial : ", f)
	}
}
