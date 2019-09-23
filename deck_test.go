package main

import (
	_ "fmt"
	"testing"
)

var (
	hands = map[string]Hand{
		"royalFlush": Hand([]Card{
			Card{suit: 0, rank: 9},
			Card{suit: 0, rank: 10},
			Card{suit: 0, rank: 11},
			Card{suit: 0, rank: 12},
			Card{suit: 0, rank: 13},
		}),

		"straightFlush": Hand([]Card{
			Card{suit: 1, rank: 8},
			Card{suit: 1, rank: 9},
			Card{suit: 1, rank: 10},
			Card{suit: 1, rank: 11},
			Card{suit: 1, rank: 12},
		}),

		"four": Hand([]Card{
			Card{suit: 0, rank: 1},
			Card{suit: 1, rank: 1},
			Card{suit: 2, rank: 1},
			Card{suit: 3, rank: 1},
			Card{suit: 0, rank: 0},
		}),

		"three": Hand([]Card{
			Card{suit: 0, rank: 0},
			Card{suit: 0, rank: 1},
			Card{suit: 1, rank: 1},
			Card{suit: 2, rank: 1},
			Card{suit: 3, rank: 2},
		}),
	}
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

func TestHand(t *testing.T) {
	deck := NewDeck()

	cards, err := deck.Draw(5)
	if err != nil {
		t.Errorf(err.Error())
	}

	hand := Hand(cards)
	oldcard := hand[1]
	err = hand.ReplaceCard(1, deck)
	if err != nil {
		t.Errorf(err.Error())
	}
	newcard := hand[1]
	if oldcard == newcard {
		t.Errorf("card not replaced")
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

	_, err = deck.Draw(1)
	if err == nil {
		t.Errorf("expected overdraw")
	}
}

// TODO: This interface sucks
func TestThree(t *testing.T) {
	cards := hands["three"].Three()
	if len(cards) != 3 {
		t.Errorf("expected 3 got %d ", len(cards))
	}
	cards = hands["four"].Three()
	if len(cards) == 0 {
		t.Errorf("expected 0 got %d ", len(cards))
	}
}
