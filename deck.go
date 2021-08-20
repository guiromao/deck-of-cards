package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Create new type of 'deck'
// which is a Slice of strings

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Diamonds", "Spades", "Cups", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// The underscore "_" will replace the need to have an index variable
	// in the for loop
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			card := value + " of " + suit
			cards = append(cards, card)
		}
	}

	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	deckBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a newDeck()
		// Option #2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	deckString := string(deckBytes)

	return deck(strings.Split(deckString, ","))
}
