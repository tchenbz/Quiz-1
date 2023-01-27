package main

import "fmt"

type Triangle struct {
	base   float64
	height float64
	hyp float64
}

func (t *Triangle) area() float64 {
	return t.base * t.height / 2
}
func (t Triangle) perimeter() float64 {
    return t.width + t.height + t.hyp
}

func main() {
	t := Triangle {
		base: 10, 
		height: 5, 
		hyp: 15
	}
	fmt.Println("area: ", t.area())
	fmt.Println("perimeter:", t.perimeter())
}
