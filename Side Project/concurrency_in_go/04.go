//? Detecting Data Race

//todo: go run -race main.go
//* inbuilt tool, to detect data races

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var n int = 0

	var wg sync.WaitGroup

	wg.Add(200)

	for i := 0; i < 100; i++ { //* Total goroutines = 200
		go func() {
			defer wg.Done()

			time.Sleep(time.Second / 10) // Simulate a complicated task

			n++
		}()

		go func() {
			defer wg.Done()

			time.Sleep(time.Second / 10) // Simulate a complicated task

			n--
		}()
	}

	wg.Wait() // Make main() wait, until all goroutines have finished executing

	fmt.Println("Final value of n: ", n) //! Should be equal to 0
}

/*
? OUTPUT :-

aryan@aryan-ZenBook-UX425JA-UX425JA:~/publisher_subscriber_golang/concurrency_in_go$ go run -race 04.go
==================
WARNING: DATA RACE
Read at 0x00c00013c008 by goroutine 8:
  main.main.func2()
      /home/aryan/publisher_subscriber_golang/concurrency_in_go/04.go:35 +0x7e

Previous write at 0x00c00013c008 by goroutine 7:
  main.main.func1()
      /home/aryan/publisher_subscriber_golang/concurrency_in_go/04.go:27 +0x90

Goroutine 8 (running) created at:
  main.main()
      /home/aryan/publisher_subscriber_golang/concurrency_in_go/04.go:30 +0x99

Goroutine 7 (finished) created at:
  main.main()
      /home/aryan/publisher_subscriber_golang/concurrency_in_go/04.go:22 +0x156
==================
Final value of n:  -1
Found 1 data race(s)
exit status 66
aryan@aryan-ZenBook-UX425JA-UX425JA:~/publisher_subscriber_golang/concurrency_in_go$
*/
