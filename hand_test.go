package main

import (
	_ "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
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

		"2pair": Hand([]Card{
			Card{suit: 0, rank: 1},
			Card{suit: 1, rank: 1},
			Card{suit: 2, rank: 2},
			Card{suit: 3, rank: 2},
			Card{suit: 0, rank: 0},
		}),

		"1pair": Hand([]Card{
			Card{suit: 0, rank: 0},
			Card{suit: 1, rank: 1},
			Card{suit: 2, rank: 2},
			Card{suit: 3, rank: 4},
			Card{suit: 0, rank: 4},
		}),

		"fullhouse": Hand([]Card{
			Card{suit: 0, rank: 0},
			Card{suit: 1, rank: 0},
			Card{suit: 2, rank: 2},
			Card{suit: 3, rank: 2},
			Card{suit: 0, rank: 2},
		}),

		"straight": Hand([]Card{
			Card{suit: 0, rank: 4},
			Card{suit: 1, rank: 3},
			Card{suit: 2, rank: 2},
			Card{suit: 3, rank: 1},
			Card{suit: 0, rank: 0},
		}),

		"flush": Hand([]Card{
			Card{suit: 1, rank: 4},
			Card{suit: 1, rank: 3},
			Card{suit: 1, rank: 2},
			Card{suit: 1, rank: 1},
			Card{suit: 1, rank: 0},
		}),
	}
)

func TestReplaceCard(t *testing.T) {
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

func TestHandes(t *testing.T) {
	if true {
		assert.Equal(t, 0, len(hands["three"].Pair()))
		assert.Equal(t, 3, len(hands["three"].Three()))
		assert.Equal(t, 0, len(hands["three"].Four()))
		assert.Equal(t, 0, len(hands["four"].Pair()))
		assert.Equal(t, 0, len(hands["four"].Three()))
		assert.Equal(t, 4, len(hands["four"].Four()))
		assert.Equal(t, 2, len(hands["1pair"].Pair()))
		assert.Equal(t, 0, len(hands["1pair"].Three()))
		assert.Equal(t, 0, len(hands["1pair"].Four()))
		assert.Equal(t, 4, len(hands["2pair"].Pair()))
		assert.True(t, hands["royalFlush"].RoyalFlush())
		assert.True(t, hands["straightFlush"].StraightFlush())
		assert.True(t, hands["fullhouse"].FullHouse())
		assert.True(t, hands["flush"].Flush())
		assert.True(t, hands["straight"].Straight())
	}
}
