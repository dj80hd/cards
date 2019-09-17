package main

import (
	"fmt"
	"testing"
)

func TestDeck(t *testing.T) {
	deck1 := NewDeck()
	deck2 := NewDeck()
	if 52 != deck1.Len() {
		t.Errorf("deck has %d", deck1.Len())
	}

	if fmt.Sprintf("%s", deck1) == fmt.Sprintf("%s", deck2) {
		t.Errorf("these decks should not be the same: %s %s", deck1, deck2)
	}
}

func TestDraw(t *testing.T) {
	deck := NewDeck()
	card := deck.Draw()
	fmt.Println(card)
	if 51 != deck.Len() {
		t.Errorf("deck has %d after draw", deck.Len())
	}
}
