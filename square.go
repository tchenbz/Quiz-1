/*package main

import "fmt"

type square struct {
	side float64
}

func (s *square) area() float64 {
	return s.side * s.side
}

func (s square) perimeter() float64 {
	return s.side + s.side + s.side + s.side
}

func main() {
	s := square {
		side: 5,
	}
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
}