package basic

import (
	"fmt"
	"math"
)

// Define interface type
type shape interface {
	area() float64
}

// Structs of shapes
type circle struct {
	radius float64
}

type rectangle struct {
	length, breadth float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func getArea(s shape) float64 {
	return s.area()
}

/*Interfaces sample interface*/
func Interfaces() {
	c := circle{radius: 5}
	fmt.Println(getArea(c))
	r := rectangle{length: 10, breadth: 20}
	fmt.Println(getArea(r))
}
