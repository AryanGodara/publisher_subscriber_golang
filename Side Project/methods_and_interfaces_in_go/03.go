//? Empty interface
package main

import "fmt"

type emptyInterface interface{}

func main() {
	var empty interface{} // A new variable of type empty interface

	empty = 5
	fmt.Println(empty)

	empty = "Go"
	fmt.Println(empty)

	empty = []int{1, 2, 3}
	fmt.Println(empty)
	// fmt.Println(len(empty))	//! We can't directly use interface, w/o type assertions
	slice, ok := empty.([]int)
	if ok != true {
		fmt.Println("There was an error in converting")
	} else {
		fmt.Println(len(slice))
	}
	// OR, if you're sure, and there's no need to check
	fmt.Println(empty.([]int))

	fmt.Printf("\n\n\n\n")

	person := make(map[string]interface{}, 0)
	//* Now, we can store any type of 'value' for 'string type' keys
	person["name"] = "Aryan"
	person["age"] = 21
	person["height"] = 181.5

	fmt.Printf("%+v\n\n", person)

	// person["age"] += 1	//! Still invalid, do TYPE ASSERTION first
}
