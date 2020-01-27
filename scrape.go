package main

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
)

const site = "https://www.monster.com/jobs/search/?q=Software-Developer&where=Australia"

type item struct {
	//Name string
	// Username string
	Title string
}

func main() {
	c := colly.NewCollector()

	messages := []item{}

	c.OnHTML(".card-content", func(e *colly.HTMLElement) {
		messages = append(messages, item{
			Title: e.ChildText("div .summary .card-header .title"),
			// ,
			// Username: e.ChildText(".account-group .username"),
			// Message:  e.ChildText(".tweet-text"),
		})
	})

	err := c.Visit(site)
	if err != nil {
		panic(err)
	}

	c.Wait()

	bs, err := json.MarshalIndent(messages, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
	fmt.Println("Number of tweets:", len(messages))
}
