package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
    d := newDeck()

    // deck should have 52 cards
    if len(d) != 52 {
        t.Errorf("Expected deck length 52, but got %v", len(d))
    }

    // first card should be Ace of Spades
    if d[0] != "Ace of Spades" {
        t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
    }

    // last card should be King of Clubs
    if d[len(d) - 1] != "King of Clubs" {
        t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d) - 1])
    }
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
    test_file_name := "_decktesting"
    os.Remove(test_file_name)

    deck := newDeck()
    deck.saveToFile(test_file_name)

    loadedDeck := newDeckFromFile(test_file_name)

    if len(loadedDeck) != 52 {
        t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
    }

    os.Remove(test_file_name)
}
