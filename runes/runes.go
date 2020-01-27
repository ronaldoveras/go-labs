package main

import "fmt"

func main() {
	str := []rune{'♪', '♬', 'a', 'A', '◍'}

	for _, v := range str {
		fmt.Printf("O caracter é %c com Código unicode %U \n", v, v)
	}
}
