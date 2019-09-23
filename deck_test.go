package main

import (
	_ "fmt"
	"testing"
)

func TestDeck(t *testing.T) {
	deck1 := NewDeck()
	deck2 := NewDeck()

	if 52 != deck1.Len() {
		t.Errorf("deck has %d", deck1.Len())
	}

	if deck1.String() == deck2.String() {
		t.Errorf("these decks should not be the same: %s %s", deck1, deck2)
	}
}

func TestOverDraw(t *testing.T) {
	deck := NewDeck()

	cards, err := deck.Draw(52)
	if err != nil {
		t.Errorf(err.Error())
	}

	if 52 != len(cards) {
		t.Errorf("draw whole desk is %d", len(cards))
	}

	_, err = deck.Draw(1)
	if err == nil {
		t.Errorf("expected overdraw")
	}
}
