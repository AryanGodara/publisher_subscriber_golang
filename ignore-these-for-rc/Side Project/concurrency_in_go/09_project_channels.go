//? Project Refactoring using waitgroups and goroutines
//todo: We want the application to check the urls, and save them in a concurrent manner
//! Not sequentially

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
)

//todo: 2. Change the function to use the channel

func checkAndSaveBody(url string, c chan string) {
	resp, err := http.Get(url)
	if err != nil {
		// fmt.Println("http.Get() failed:", err)
		s := fmt.Sprintf("%s is DOWN!\n", url) // Sprint formats using the default formats for its operands and returns the resulting string. Spaces are added between operands when neither is a string
		s += fmt.Sprintf("Error: %v\n", err)

		//* This is a blocking call, so this goroutine will wait for the main goroutine to
		//* receive it on the other part of the channel
		c <- s // sending into the channel
	} else {
		defer resp.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d\n", url, resp.StatusCode)

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

			s += fmt.Sprintf("Writing response body to %s\n", file)

			//todo: this function will from ioutil will handle: creating, opening, writing slice of bytes, and closing the file

			err = ioutil.WriteFile(file, bodyBytes, 0664) // 0664 = file permission, google
			if err != nil {
				// log.Fatal("Couldn't write to file", err) //Fatal is equivalent to Print() followed by a call to os.Exit(1).
				s += "Error writing file\n"
				c <- s
			}
		}

		s += fmt.Sprintf("%s is UP\n", url)
		c <- s
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

	for _, url := range urls {
		go checkAndSaveBody(url, c)
		fmt.Println(strings.Repeat("#", 20))
	}

	fmt.Println("No. of goroutines: ", runtime.NumGoroutine())

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}
