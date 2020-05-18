package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of deck, this deck is a slice of strings
// This type is an approximation to the oop concept in golang
// which is not an object oriented language
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

// saveToFile function saves into the hard drive a deck instance
// This is made using the ioutil library WriteFile function.
// To do this the file name is taken as parameter, the byte slice
// is transformed directly from the deck slice to string conversion
// (using the previous toString function) and finally the permissions
// are set to 0666. In case of error an error is returned
func (d deck) saveToFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
}

// newDeckFromFile function recovers from file an existing deck
// of cards. This returns a slice of bytes and (in case of happening)
// an error (nil when no error happened)
// If the file has been read properly theres a chance this file is
// empty. In this case the execution stops
func newDeckFromFile(fileName string) deck {
	bs, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

// shuffle function performs a deck shuffleling using pseudo
// random function rand. Using this number changes position
// of elements in deck. Being a pseudo random number generates
// an issue which is the chance of getting always the same result
// increases. To avoid this a new source is adde to the rand
// Unix time in nanoseconds is used as new seed
func (d deck) shuffle() {
	for i := range d {

		source := rand.NewSource(time.Now().UnixNano())
		r := rand.New(source)
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
