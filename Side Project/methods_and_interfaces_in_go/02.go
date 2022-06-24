//? Interface embedding

package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type object interface {
	volume() float64
}

// geometry is embedding shape and object interfaces
//* When we embed an interface to another, we add all the methods from the embedded
//* interface, to the new interface
type geometry interface {
	shape
	object
	getColor() string
}

type cube struct {
	edge  float64
	color string
} //* cube type implements both "shape" and "object" interfaces

func (c cube) area() float64 {
	return math.Pow(c.edge, 2)
}

func (c cube) volume() float64 {
	return math.Pow(c.edge, 3)
}

func measure(g geometry) (float64, float64) {
	a := g.area()
	v := g.volume()

	return a, v
}

//? We also need to define a 'getColor()' method for cube type, in order for it to
//? implement the 'geometry' interface
func (c cube) getColor() string {
	return c.color
}

func main() {
	c := cube{3, "red"}
	a, b := measure(c)
	fmt.Println(a, b)
}

/*
! The following results in compile-time error (Circular embedding)

type I1 interface {
	I2
	m1()
}

type I2 interface {
	I3
	m2()
}

type I3 interface {
	I1
	m3()
}
*/
