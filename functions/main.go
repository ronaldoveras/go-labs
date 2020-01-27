package main

import (
	"fmt"
)

func main() {
	d := soma(4, 6)
	fmt.Println(d)
	p := 10

	r := double(p)
	fmt.Printf("Valor original: %d \n", p)
	fmt.Printf("Valor dobrado: %d \n", r)

	q := doublePointer(&p)
	fmt.Printf("Valor original (pointer): %d \n", p)
	fmt.Printf("Valor dobrado (pointer): %d \n", q)

}

func soma(a int32, b int32) (z int32) {
	z = a + b
	return
}

func doublePointer(a *int) int {
	*a = 2 * (*a)
	return *a
}

func double(a int) int {
	a = 2 * (a)
	return a
}
