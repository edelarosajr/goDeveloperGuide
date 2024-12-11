package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card to be Ace of Clubs, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected first card to be King of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFile := "_decktesting"
	os.Remove(testFile)

	d := newDeck()
	d.saveToFile(testFile)

	loadedDeck := newDeckFromFile(testFile)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, but got %v", len(loadedDeck))
	}

	os.Remove(testFile)
}
