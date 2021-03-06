package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { //t represent our test handles
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck lenght 16, but got %v", len(d)) //%v its the equivalent of {} on formated python strings
	}

	if d[0] != "As of Spades" {
		t.Errorf("Expected first card of As of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of CLubs, but got %v", d[len(d)-1])
	}

}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckteseting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
