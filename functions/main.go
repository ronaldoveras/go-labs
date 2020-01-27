package main

import (
	"fmt"
)

func main() {
	d := soma(4, 6)
	fmt.Println(d)

}

func soma(a int32, b int32) (z int32) {
	z = a + b
	return
}
