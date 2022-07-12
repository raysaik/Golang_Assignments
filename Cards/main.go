package main

import (
	"fmt"
)

var deckSize int

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of Spades" //only used for new variable not already declared variables
	card = newCard()
	fmt.Println(card)

	cards := newDeck()
	fmt.Println(cards)

	cards.print()

	cards = append(cards, "Squeen of Scades") //Immutable operation- create a new variable and assignes to cards

	// for card := range cards {
	// 	fmt.Println(card)
	// }
	deckSize = 25
	fmt.Println(deckSize)
}

func newCard() string {
	return "Five Of Spades"
}
