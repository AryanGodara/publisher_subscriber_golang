//? WaitGroups

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("\nf1 (goroutine) execution started")

	defer fmt.Println("f1 execution finished\n")

	for i := 0; i < 3; i++ {
		fmt.Println("f1, i = ", i)
		time.Sleep(time.Second) // fake simulation of a long task
	}
	//* we call wg.done() in each goroutine to indicate to the waitgroup that the
	//* goroutine has finished executing

	wg.Done()
}

func f2(wg *sync.WaitGroup) {
	fmt.Println("\nf2 (goroutine) execution started")

	defer fmt.Println("f2 execution finished\n")

	for i := 5; i < 8; i++ {
		fmt.Println("f2, i = ", i)
		time.Sleep(time.Second) // fake simulation of a long task
	}

	wg.Done()
	// (*wg).Done() -> Same thing, as go handles it internally
}

func main() {

	var wg sync.WaitGroup

	wg.Add(2) // 2 = no. of goroutines to wait for. Once for go f1(), once for go(f2)

	fmt.Println("main execution started")
	defer fmt.Println("Main execution stopped")

	fmt.Println("No. of CPUs: ", runtime.NumCPU())
	fmt.Println("No. of Goroutines: ", runtime.NumGoroutine())
	fmt.Println()
	fmt.Println("OS: ", runtime.GOOS)
	fmt.Println("Arch: ", runtime.GOARCH)
	fmt.Println()
	fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0))
	fmt.Println()

	go f1(&wg)
	fmt.Println("No. of Goroutines after go f1(): ", runtime.NumGoroutine())

	f2(&wg)
	fmt.Println("No. of Goroutines after f2(): ", runtime.NumGoroutine())

	// go f2(&wg)
	// fmt.Println("No. of Goroutines after go f2(): ", runtime.NumGoroutine())

	wg.Wait() // Wait until all goroutines have stopped
}

/*
//! OUTPUT :_

aryan@aryan-ZenBook-UX425JA-UX425JA:~/publisher_subscriber_golang/concurrency_in_go$ go run 02.go
main execution started
No. of CPUs:  8
No. of Goroutines:  1

OS:  linux
Arch:  amd64

GOMAXPROCS:  8

No. of Goroutines after go f1():  2

f2 (goroutine) execution started
f2, i =  5

f1 (goroutine) execution started
f1, i =  0
f1, i =  1
f2, i =  6
f2, i =  7
f1, i =  2
f2 execution finished

No. of Goroutines after f2():  2
Main execution stopped
aryan@aryan-ZenBook-UX425JA-UX425JA:~/publisher_subscriber_golang/concurrency_in_go$
*/
