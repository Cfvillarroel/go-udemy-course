package main

import "fmt"

// Create a new type of deckType
type deckType [] string

func newDeck() deckType {
	cards := deckType{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King" }

	for _, cardSuit := range cardSuits{ //the underscores mean that you know that there is a varible but you are not going to use it
		for _, cardValue := range cardValues{
			cards = append(cards, cardValue + " of " + cardSuit)
		}
	}

	return cards
}

func (deck deckType) print() {
	for i, card := range deck {
		fmt.Println(i, card)
	}
}

func deal(deck deckType, handSize int) (deckType, deckType) {
	return deck[:handSize], deck[handSize:]
} 