//? Project Refactoring using waitgroups and goroutines
//todo: We want the application to check the urls, and save them in a concurrent manner
//! Not sequentially

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("http.Get() failed:", err)
		fmt.Printf("%s is DOWN!\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d\n", url, resp.StatusCode)

		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("Coudn't read file", err)
			}

			file := strings.Split(url, "//")[1]
			// http://www.google.com
			//todo: Split by // -> []{"http:", "www.google.com"} -> [1:], ignore http:

			file += ".txt" //* file -> google.com.txt
			file = "./project_3_files/" + file

			fmt.Printf("Writing response body to %s\n", file)

			//todo: this function will from ioutil will handle: creating, opening, writing slice of bytes, and closing the file

			err = ioutil.WriteFile(file, bodyBytes, 0664) // 0664 = file permission, google
			if err != nil {
				log.Fatal("Couldn't write to file", err)
			}
		}
	}

	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	urls := []string{
		"https://golang.org",
		"https://golanggg.orrg",
		"https://www.google.com/randfilename.html",
		"https://www.google.com",
		"https://www.medium.com"}

	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg)
		fmt.Println(strings.Repeat("#", 20))
	}

	fmt.Println("No of Goroutines: ", runtime.NumGoroutine())
	//* Should be = 6 (1 main goroutine, and 5 for len(urls))

	wg.Wait()
	//! Don't close main(), until all the goroutines inside main() have finished executing
}
