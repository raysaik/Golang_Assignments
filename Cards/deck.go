package main

import (
	"fmt"
	"strings"
)

//create a new type 'deck' which is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func NewDeck() deck {
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

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) ToString() string {
	return strings.Join([]string(d), ",")
}
