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
	cards := deck.Draw(1)
	fmt.Println(cards)
	if 51 != deck.Len() {
		t.Errorf("deck has %d after draw", deck.Len())
	}
}

func TestHand(t *testing.T) {
	deck := NewDeck()
	hand := deck.Draw(5)
	fmt.Println(hand)
	if 47 != deck.Len() {
		t.Errorf("deck has %d after draw", deck.Len())
	}
}
