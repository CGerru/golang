package main

import (
	"os"
	"testing"
)

// TestNewDeck is a test function which point is
// to check the behavior of the newDeck() function
// Every test funcion names in Go MUST start with
// a capital letter. Also testing.T must enter as
// parameter
func TestNewDeck(t *testing.T) {
	d := newDeck()

	// Check the length of the deck
	if len(d) != 28 {
		t.Errorf("Expected 28 length, got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("First card expected Ace of Spades, got %v", d[0])
	}

	if d[len(d)-1] != "Seven of Clubs" {
		t.Errorf("Last card expected Seven of Clubs, got %v", d[len(d)-1])
	}
}

// TestSaveToDeckAndNewDeckFromFile function checks the behavior of
// the saving getting decks into the hard drive. First step is to delete
// the file used to perform test, in case previous test executions has not
// gone well.
func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 28 {
		t.Errorf("Expected 28 length, got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
