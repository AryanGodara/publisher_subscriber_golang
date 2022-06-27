package main

import (
	"fmt"
	"sync"
)

func main() {
	myCh := make(chan int)
	wg := &sync.WaitGroup{}

	//! DEADLOCK
	// myCh <- 5
	// fmt.Println(<-myCh)

	wg.Add(2)

	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myCh)
		wg.Done()
	}(myCh, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		myCh <- 5
		wg.Done()
	}(myCh, wg)

	wg.Wait()
}
