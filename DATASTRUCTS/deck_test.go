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

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be 'Ace of Spades', but got '%v'", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be 'King of Clubs', but got '%v'", d[len(d)-1])
	}
}

func TestSaveToDeckAndLoadFromFile(t *testing.T) {
	file := "_decktesting"
	os.Remove(file)

	d := newDeck()
	d.saveToFile(file)

	loadedDeck := loadFromFile(file)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove(file)
}
