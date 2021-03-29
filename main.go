package main

func main(){
	 cards := newDeck()
/*
	hand, reaminingCards := deal(cards, 5)

	hand.print()
	reaminingCards.print();
 */
	cards.saveToFile("my_cards")

}