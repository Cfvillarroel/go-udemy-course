package main

import (
	"fmt" 
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

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

func (deck deckType) toString() string {
	return strings.Join([]string(deck), ",")
}

func (deck deckType) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(deck.toString()), 0666)
}

func newDeckFromFile(filename string) deckType {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option 1 - log the error and return call no newDeck()
		// Option 2 - log the error en entirely quit the program
		fmt.Println("Error:", err)
		os.Exit(1)	
	}

	cardString := strings.Split(string(bs), ",")
	return deckType(cardString)
}

func (deck deckType) shuffle(){
	source := rand.NewSource(time.Now().UnixNano()) // generate random int64
	r := rand.New(source) // generate the rand seed

	for i := range deck { // shuffle the deck
		newPosition := r.Intn(len(deck) - 1)

		deck[i], deck[newPosition] = deck[newPosition], deck[i]
	}
}