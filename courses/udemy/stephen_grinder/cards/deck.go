package main

import (
	"fmt"
	"strings"
)

// Create a new type of deck, this deck is a slice of strings
// This type is an approximation to the oop concept in golang
// wich is not an object oriented language
type deck []string

//newDeck creates a new deck out of nothing
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven"}

	// Using the underscore '_' we acknowledge that, instead of
	// declaring a variable (in this cas i and j), we ignore this.
	// Declaring the indexes specifically would throw the "declared
	// variable but not used error"
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// print prints the cards in the deck
// This is a receiver function that enables any deck type
// access to this, so can call this function on itself
// d is the instance of the deck variable we are working with
// By convetion the receiver variable name is named by
// one or two abbreviation
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// deal function takes two parameters one instance of deck type
// and an integer of the number of cards the deck has by size.
// This function is returning two values
//
// One deck type with the result of the deal and other deck type
// with the remaining cards on deck
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// toString function returns a deck type onto a string. To do this
// the function initially makes the conversion of the type deck onto
// a slice []string (witch is essentially) and includes all its values
// in a new string. This process is performed with the strings library
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
