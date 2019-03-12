package main

import (
	"fmt"
)

func main() {
        v:=1
	fmt.Println(incr(&v))
	fmt.Println(incr(&v))
	fmt.Println(v)
}

func incr(p *int) int{
   *p++
   return *p
}
