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

func TestHand(t *testing.T) {
	deck := NewDeck()

	cards, err := deck.Draw(5)
	if err != nil {
		t.Errorf(err.Error())
	}

	hand := Hand(cards)
	err = hand.NewCard(1, deck)
	if err != nil {
		t.Errorf(err.Error())
	}

	if 46 != deck.Len() {
		t.Errorf("deck has %d after draw", deck.Len())
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

	cards, err = deck.Draw(1)
	if err == nil {
		t.Errorf("expected overdraw")
	}

}
