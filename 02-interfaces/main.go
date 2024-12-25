package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func main() {
	c1 := circle{
		radius: 10,
	}
	s1 := square{
		length: 5,
		width:  5,
	}
	fmt.Println("Area of square s1: ", s1.area())
	fmt.Println("Area of circle c1: ", c1.area())

}
