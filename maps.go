package main

import (
	"fmt"
	"strconv"
)

func main() {
	m := make(map[int]string)
	for i:=1; i<15; i++ {
		m[i] = "Olá " + strconv.Itoa(i)
	}
	for _, val := range m {
		fmt.Println(val)
	}
}

