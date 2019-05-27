package main

import (
	"fmt"
)

func main() {
	const s = "\x41"
	fmt.Println(s)
	const b = "Hello World"
	for i,r := range b {
		fmt.Printf("%d %#U\n", i, r)
	}
}
