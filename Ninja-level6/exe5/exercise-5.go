// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// 	circle area= π r 2
// 	square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

package main

import (
	"fmt"
	"math"
)

type SQUARE struct {
	length float64
}

type CIRCLE struct {
	radius float64	
}

func (s SQUARE) area() float64 {
	return s.length * s.length
}

func (c CIRCLE) area() float64 {
	return c.radius * c.radius * math.Pi
}

// Create an interface
type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	circle := CIRCLE{
		radius: 7.678,
	}
	square := SQUARE{
		length: 15,
	}

	info(circle)
	info(square)
}