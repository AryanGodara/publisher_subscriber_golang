package main

import (
	"fmt"
	"math"
)

type Shapes interface { // Both cirle and rectangle structs implement this interface
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) volume() float64 {
	return (4 * math.Pi * math.Pow(c.radius, 3)) / 3
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func printShape(s Shapes) {
	fmt.Printf("Shape Features 1: %#v\n", s)
	fmt.Printf("Shape Features 2: %+v\n", s)
	fmt.Printf("Shape Features 3: %v\n", s)
	fmt.Println("Area : ", s.area())
	fmt.Println("Perimeter : ", s.perimeter())
}

func main() {
	c := circle{2}
	r := rectangle{2, 3}

	// fmt.Println(c.area())
	// fmt.Println(c.perimeter())
	// fmt.Println(r.area())
	// fmt.Println(r.perimeter())
	printShape(c)
	fmt.Println()
	printShape(r)

	var s Shapes = circle{2.5}
	// fmt.Println(s.volume)
	//! Can't use volume method, since that is only a part of 'circle' type, not a part of the 'Shapes' interface
	ball, ok := s.(circle) //* TYPE ASSERTION
	if !ok {
		fmt.Println("Assertion wasn't successful")
	} else {
		fmt.Println(ball.volume()) // Now, we can do this
	}

	//? Type Switches
	switch value := s.(type) { // Switch statement for 'value' variable
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)
	}
}
