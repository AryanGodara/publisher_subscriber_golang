//? Project Refactoring using waitgroups and goroutines
//todo: We want the application to check the urls, and save them in a concurrent manner
//! Not sequentially

package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

//todo: 2. Change the function to use the channel

func checkURL(url string, c chan string) {
	resp, err := http.Get(url)

	if err != nil {
		s := fmt.Sprintf("%s is DOWN!\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		fmt.Println(s)

		//* This is a blocking call, so this goroutine will wait for the main goroutine to
		//* receive it on the other part of the channel
		c <- url // sending into the channel
	} else {
		s := fmt.Sprintf("%s -> Status Code: %d\n", url, resp.StatusCode)
		s += fmt.Sprintf("%s is UP\n", url)
		fmt.Println(s)

		c <- url
	}

}

func main() {
	urls := []string{
		"https://golang.org",
		"https://golanggg.orrg",
		"https://www.google.com/randfilename.html",
		"https://www.google.com",
		"https://www.medium.com"}

	//todo: 1.Create and initialize a goroutine, of type 'string'
	c := make(chan string)

	for _, url := range urls { // Enter each url to the goroutine
		go checkURL(url, c)
	}

	fmt.Println("No. of goroutines: ", runtime.NumGoroutine())

	// for i := 0; i < len(urls); i++ {
	// 	fmt.Println(<-c)
	// }

	// for { // INFINITE LOOP
	// 	go checkURL(<-c, c)
	// 	// url is received from the previous goroutine, so the checkURL will infinitely
	// 	// keep checking the urls, without stopping
	// 	fmt.Println(strings.Repeat("#", 30))

	// 	time.Sleep(time.Second)
	// }

	//? Another way of doing this:
	// for url := range c {
	// 	time.Sleep(time.Second)
	// 	go checkURL(url, c)
	// 	fmt.Println(strings.Repeat("#", 30))
	// }

	//? Another way, using anon funcs
	for url := range c {
		go func(u string) {
			time.Sleep(time.Second)
			checkURL(u, c)
		}(url)
	}
}
