package main

func main() {
	cards := deck{newCard("As de ouro"), newCard("Reis de Espada")}
	cards = append(cards, "2 de paus")

	cards.print()
}

func newCard(name string) string {
	return name
}
