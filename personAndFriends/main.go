package main

import (
	"fmt"

	"github.com/ronaldoveras/go-labs/personAndFriends/model"
)

func main() {
	person := new(model.Person)
	person.Friends = new(model.Friends)
	person.Name = "Claúdia"
	person.Friends.Add("Bira")
	person.Friends.Add("Karol")
	person.Friends.Add("Katy")

	fmt.Printf("Amigos de %s são: \n", person.Name)
	for _, p := range person.Friends.Data {
		fmt.Println(p)
	}
}
