//? mutex : mutual exclusion -> Prevents race conditions

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

	//todo 1. create the mutex variable
	var m sync.Mutex

	for i := 0; i < 100; i++ { //* Total goroutines = 200
		go func() {
			defer wg.Done()

			time.Sleep(time.Second / 10) // Simulate a complicated task

			// The lock method of mutex variable blocks the access to the variable, until
			// the unlock method is called.
			//todo 2. Block the access to n, until this goroutine finishes

			m.Lock()
			n++
			m.Unlock()

			//* Code b/w Lock() and Unlock() is executed by ONLY 1 goroutine at a time
		}()

		go func() {
			time.Sleep(time.Second / 10) // Simulate a complicated task

			m.Lock()
			defer m.Unlock() // Will be called just before the function exits
			//? remember, defer statements are called in stack type, reverse order
			n--

			wg.Done()

		}()
	}

	wg.Wait() // Make main() wait, until all goroutines have finished executing

	fmt.Println("Final value of n: ", n) //! Should be equal to 0
}
