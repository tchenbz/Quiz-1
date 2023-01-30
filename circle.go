/*package main

import "fmt"

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return (22 / 7) * (c.radius * c.radius)
}

func (c circle) perimeter() float64 {
	return (2) * (22 / 7) * (c.radius)
}

func main() {
	c := circle{
		radius: 5,
	}
	fmt.Println("Area: ", c.area())
	fmt.Println("Perimeter: ", c.perimeter)
}