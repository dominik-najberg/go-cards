package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %d", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades but got %v", d[0])
	}

	if d[len(d) -1] != "One of Clubs" {
		t.Errorf("Expected last card of Ace of Spades but got %v", d[len(d) -1])
	}
}

func TestSaveToDiskAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")

	loadedDeck := newDeckFromFile("_deckTesting")

	if len(deck) != len(loadedDeck) {
		t.Errorf("The number of cards is wrong");
	}

	for i, card := range deck {
		if loadedDeck[i] != card {
			t.Errorf("Card %v is different from %v, at position: %d", card, loadedDeck[i], i)
		}
	}

	os.Remove("_deckTesting")
}