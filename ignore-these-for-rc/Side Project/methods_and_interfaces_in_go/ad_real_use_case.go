//? Reducing boilerplate code
/*
Imagine we have a "Customer" struct. In one part of our codebase, we want to write the
customer info to a "bytes.Buffer", and in another, to an "os.File" on disk. But in both
cases, we want to serialize the customer struct to JSON first.
* We can use interfaces.

Go has an io.Writer interface type, which looks like this

type Writer interface {
	Write(p []byte) (n int, err error)
}

Both bytes.Buffer and os.File satisfy this interface, due to them having the
bytes.Buffer.Write() and os.File.Write() method, respectively
*/

package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
)

type Customer struct {
	Name string
	Age  int
}

// Implement a WriteJSON method that takes an io.Writer as the parameter.
// It marshals the customer struct to JSON, and if the marshal worked successfully, then
// calls the relevel io.Writer's Write method
func (c *Customer) WriteJSON(w io.Writer) error {
	js, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = w.Write(js)

	return nil // No error
}

func main() {
	c := &Customer{Name: "Aryan", Age: 21}

	// We can then call the WriteJSON method using a buffer
	var buf bytes.Buffer
	err := c.WriteJSON(&buf)
	if err != nil {
		log.Fatal(err)
	}

	// Or using a file
	f, err := os.Create("/tmp/customer")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	err = c.WriteJSON(f)
	if err != nil {
		log.Fatal(err)
	}
}
