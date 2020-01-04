package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Movie struct {
	Title  string
	Year   int      `json:"released"`
	Color  bool     `json:"color,omitempy"`
	Actors []string `json:"actors"`
}

func main() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steve McQuenn", "Jacqueline Bisset"}},
	}

	data, err := json.MarshalIndent(movies, "", " ")
	if err != nil {
		log.Fatalf("Falha")
	}
	fmt.Printf("%s\n", data)

}
