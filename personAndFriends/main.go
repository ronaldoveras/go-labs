package main

import "fmt"

type friends struct {
	data []string
}

type person struct {
	name    string
	friends *friends
}

func (f *friends) Add(name string) {
	f.data = append(f.data, name)
}

func (f *friends) Remove(name string) {

}

func main() {
	person := new(person)
	person.friends = new(friends)
	person.name = "Claúdia"
	person.friends.Add("Bira")
	person.friends.Add("Karol")

	fmt.Printf("Amigos de %s são: \n", person.name)
	for _, p := range person.friends.data {
		fmt.Println(p)
	}
}
