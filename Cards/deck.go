package main

import "fmt"

//create a new type 'deck' which is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "2", "3", "4"}

	for _, cs := range cardSuites {
		for _, cv := range cardValues {
			cards = append(cards, cv+" of "+cs)
		}
	}
	return cards
}
