package main

import (
	"fmt"
)

type t interface {
	area()
	volume()
}

type item struct {
	radius float64
	height float64
}

type fake struct{}

func (m item) area() float64 {

	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}

func (m item) volume() float64 {

	return 3.14 * m.radius * m.radius * m.height
}

func (m fake) area() float64 {
	return -1
}

func main() {
	var t item
	t = item{10, 14}
	fmt.Println("Área :", t.area())
	fmt.Println("Volume:", t.volume())

	var q fake
	q = fake{}
	fmt.Println("Área :", q.area())
}
