package main

import "fmt"

var deckSize int

func main() {
	//var card string = "Ace of Spades"
	card := "Ace of Spades" //only used for new variable not already declared variables
	card = newCard()
	fmt.Println(card)

	cards := NewDeck()
	fmt.Println(cards)

	cards.print()

	cards = append(cards, "Squeen of Scades") //Immutable operation- create a new variable and assignes to cards

	deckSize = 25
	fmt.Println(deckSize)

	//Deal Cards Code
	hand, remainingcards := deal(cards, 5)
	hand.print()
	fmt.Println("BREAK")
	remainingcards.print()

	//string operations
	cardsTwo := NewDeck()
	fmt.Println(cardsTwo.ToString())

	//saving to file
	cardsTwo.SaveToDrive("abcd")

	//ReadFromFile
	cardsThree := GetCardsFromFile("ajbcd")
	fmt.Println("New Cards:", cardsThree)
}

func newCard() string {
	return "Five Of Spades"
}
