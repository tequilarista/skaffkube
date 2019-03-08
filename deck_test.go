package main

import (
	"os"
	"testing"
)

// testing.T is type used by the Go testing runner
func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52.  Actual length: %v", len(d))
	}
}

func TestFirstCard(t *testing.T) {
	d := newDeck()

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades.  Actual cards: %v", d[0])
	}
}

func TestLastCard(t *testing.T) {
	d := newDeck()

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected King of Clubs.  Actual cards: %v", d[len(d)-1])
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	// make sure env is clean in case something previously crashed
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52.  Actual length: %v", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
