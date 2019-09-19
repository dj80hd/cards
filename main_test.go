package main

import (
	_ "fmt"
	"sort"
	"testing"
)

var (
	fourOfAKind = Hand([]Card{
		Card{suit: 0, rank: 1},
		Card{suit: 1, rank: 1},
		Card{suit: 2, rank: 1},
		Card{suit: 3, rank: 1},
		Card{suit: 0, rank: 0},
	})

	threeOfAKind = Hand([]Card{
		Card{suit: 0, rank: 0},
		Card{suit: 0, rank: 1},
		Card{suit: 1, rank: 1},
		Card{suit: 2, rank: 1},
		Card{suit: 3, rank: 2},
	})

	fourStrait = Hand([]Card{
		Card{suit: 3, rank: 3},
		Card{suit: 1, rank: 4},
		Card{suit: 1, rank: 3},
		Card{suit: 1, rank: 2},
		Card{suit: 1, rank: 1},
	})

	fiveStrait = Hand([]Card{
		Card{suit: 1, rank: 0},
		Card{suit: 1, rank: 1},
		Card{suit: 1, rank: 2},
		Card{suit: 1, rank: 3},
		Card{suit: 1, rank: 4},
	})
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
	err = hand.ReplaceCard(1, deck)
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

	_, err = deck.Draw(1)
	if err == nil {
		t.Errorf("expected overdraw")
	}
}

func TestSortHand(t *testing.T) {
	fourStrait := fourStrait
	sort.Sort(fourStrait)

	if 1 != fourStrait[0].suit || 1 != fourStrait[0].rank {
		t.Errorf("sort failed %v", fourStrait)
	}
}

func TestThree(t *testing.T) {
	if 1 != threeOfAKind.Three() {
		t.Errorf("threeOfAKind")
	}
	if 1 != fourOfAKind.Three() {
		t.Errorf("threeOfAKind")
	}
	if -1 != fourStrait.Three() {
		t.Errorf("threeOfAKind")
	}
}

func TestStrait4(t *testing.T) {
	for _, c := range []struct {
		hand Hand
		suit int
		rank int
		desc string
	}{
		{hand: threeOfAKind, suit: -1, rank: -1, desc: "t1"},
		{hand: fourOfAKind, suit: -1, rank: -1, desc: "t2"},
		{hand: fourStrait, suit: 1, rank: 1, desc: "t3"},
		{hand: fiveStrait, suit: 1, rank: 0, desc: "t4"},
	} {
		suit, rank := c.hand.Strait4()
		if c.rank != rank {
			t.Errorf("wrong rank %d expect %d %s", rank, c.rank, c.desc)
		}
		if c.suit != suit {
			t.Errorf("wrong suit %d expect %d  %s", suit, c.suit, c.desc)
		}
	}
}

func TestStrait5(t *testing.T) {
	for _, c := range []struct {
		hand Hand
		suit int
		rank int
		desc string
	}{
		{hand: threeOfAKind, suit: -1, rank: -1, desc: "t1"},
		{hand: fourOfAKind, suit: -1, rank: -1, desc: "t2"},
		{hand: fourStrait, suit: -1, rank: -1, desc: "t3"},
		{hand: fiveStrait, suit: 1, rank: 0, desc: "t4"},
	} {
		suit, rank := c.hand.Strait5()
		if c.rank != rank {
			t.Errorf("wrong rank %d expect %d %s", rank, c.rank, c.desc)
		}
		if c.suit != suit {
			t.Errorf("wrong suit %d expect %d  %s", suit, c.suit, c.desc)
		}
	}
}
