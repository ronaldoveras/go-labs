package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamond", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suite+" of "+value)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
