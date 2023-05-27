package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected a legth of 16 but got: %d ", len(d))
	}
	if d[len(d)-1] != "Ace of Spades" {
		t.Errorf("Expected an ace of spades but got: %v ", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected a four of clubs: %v ", d[15])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Error:%v", len(loadedDeck))
	}
	os.Remove("_decktesting")

}
