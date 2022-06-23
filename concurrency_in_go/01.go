//? Spawning goroutines, the go keyword

package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 (goroutine) execution started")

	defer fmt.Println("f1 execution finished")

	for i := 0; i < 3; i++ {
		fmt.Println("f1, i = ", i)
	}
}

func f2() {
	fmt.Println("f2 (goroutine) execution started")

	defer fmt.Println("f2 execution finished")

	for i := 5; i < 8; i++ {
		fmt.Println("f, i = ", i)
	}
}

func main() {
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

	go f1()
	fmt.Println("No. of Goroutines after go f1(): ", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)

	f2()
	fmt.Println("No. of Goroutines after f2(): ", runtime.NumGoroutine())

	go f2()
	fmt.Println("No. of Goroutines after go f2(): ", runtime.NumGoroutine())

	time.Sleep(time.Second * 2)
}
