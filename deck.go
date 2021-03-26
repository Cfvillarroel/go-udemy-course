package main

import "fmt"

// Create a new type of deck
type deckType [] string

func (deck deckType) print() {
	for i, card := range deck {
		fmt.Println(i, card)
	}
}