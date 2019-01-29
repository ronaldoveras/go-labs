package main

import "fmt"

func main() {
	card := "As de espadas"
	card = newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Dois de reis"
}
