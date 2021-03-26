package main

import "fmt"

// Create a new type of deck
type deck [] string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits{ //the underscores mean that you know that there is a varible but you are not going to use it
		for _, cardValue := range cardValues{
			cards = append(cards, cardValue + " of " + cardSuit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}