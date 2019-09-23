package main

import (
	"fmt"
	"sort"
)

/**
*
* RoyalFlush bool
* StraitFlush() bool
* Four() []int
* FullHouse() bool
* Flush() bool
* Straight() bool
* Three() []int
* Pair() []int
 */

type Hand []Card

func (a Hand) Len() int      { return len(a) }
func (a Hand) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Hand) Less(i, j int) bool {
	if a[i].suit == a[j].suit {
		return a[i].rank < a[j].rank
	}
	return a[i].suit < a[j].suit
}

func (h Hand) RoyalFlush() bool {
	sort.Sort(h)                               // TODO: ensure hand is always sorted to avoid this
	return h.StraightFlush() && h[0].rank == 9 //jack
}

func (h Hand) StraightFlush() bool {
	sort.Sort(h) // TODO: ensure hand is always sorted to avoid this
	suit := h[0].suit
	rank := h[0].rank
	for i := 1; i < len(h); i++ {
		if h[i].suit != suit || h[i].rank != rank+1 {
			return false
		}
		suit = h[i].suit
		rank = h[i].rank
	}
	return true
}

func (h Hand) Four() []Card {
	sort.Sort(h) // TODO: ensure hand is always sorted to avoid this

	cards := []Card{}

	for i := 0; i < len(h)-3; i++ {
		if h[i].rank == h[i+1].rank && h[i].rank == h[i+2].rank && h[i].rank == h[i+3].rank {
			cards = append(cards, h[i], h[i+1], h[i+2], h[i+3])
		}
	}

	return cards
}

func (h Hand) FullHouse() bool {
	// TODO: this is different for 5 and 7 card hands ?
	return len(h.Pair()) == 1 && len(h.Three()) == 1
}

func (h Hand) Flush() bool {
	suit := h[0].suit
	for i := 1; i < len(h); i++ {
		if suit != h[i].suit {
			return false
		}
	}
	return true
}

func (h Hand) Straight() bool {
	sort.Sort(h) // TODO: ensure hand is always sorted to avoid this
	rank := h[0].rank
	for i := 1; i < len(h); i++ {
		if h[i].rank != rank+i {
			return false
		}
	}
	return true
}

func (h Hand) Three() []Card {
	sort.Sort(h) // TODO: ensure hand is always sorted to avoid this

	cards := []Card{}

	for i := 0; i < len(h)-2; i++ {
		if h[i].rank == h[i+1].rank && h[i].rank == h[i+2].rank {
			// don't count 4 of a kind, etc.
			if i == len(h)-3 || h[i].rank != h[i+3].rank {
				cards = append(cards, h[i], h[i+1], h[i+2])
			}
		}
	}

	return cards
}

func (h Hand) Pair() []Card {
	sort.Sort(h) // TODO: ensure hand is always sorted to avoid this

	cards := []Card{}

	for i := 0; i < len(h)-1; i++ {
		if h[i].rank == h[i+1].rank {
			if i == len(h)-2 || h[i].rank != h[i+2].rank {
				cards = append(cards, h[i], h[i+1])
			}
		}
	}

	return cards
}

// TODO: Score a hand for comparison
func (h Hand) Score() int {
	return 0
}

// ReplaceCard replaces the nth card with a new one from the deck
func (h Hand) ReplaceCard(n int, d *Deck) error {
	if n >= len(h) {
		return fmt.Errorf("index %d >= len %d", n, len(h))
	}

	cards, err := d.Draw(1)
	if err != nil {
		return err
	}
	h[n] = cards[0]

	return nil
}