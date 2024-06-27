package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected the length of 16 but got %v", len(d))
	}

	if d[0] != "Ace of Spade" {
		t.Errorf("Expected Ace of spades but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Diamond" {
		t.Errorf("Expected Four of Diamond but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	deck := newDeck()

	deck.savetofile("_decktesting")

	loaddeddeck := newDeckfromfile("_decktesting")

	if len(loaddeddeck) != 16 {
		t.Errorf("Expected 16 cards but got %v", len(loaddeddeck))
	}

	os.Remove("_decktesting")
}
