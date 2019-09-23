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

// Three returns a slice of indexes where 3 of a kind is found.
func (h Hand) Three() []int {
	sort.Sort(h) // TODO: ensure hand is always sorted to avoid this

	ret := []int{}

	for i := 0; i < len(h)-2; i++ {
		if h[i].rank == h[i+1].rank && h[i].rank == h[i+2].rank {
			ret = append(ret, i)
		}
	}

	return ret
}

//Strait4 returns the suit and rank of a straight. -1 for none
func (h Hand) Strait4() (int, int) {
	sort.Sort(h)
	for start := 0; start < len(h)-3; start++ {
		if h[start].suit == h[start+1].suit &&
			h[start].suit == h[start+2].suit &&
			h[start].suit == h[start+3].suit &&
			h[start].rank+1 == h[start+1].rank &&
			h[start].rank+2 == h[start+2].rank &&
			h[start].rank+3 == h[start+3].rank {
			return h[start].suit, h[start].rank
		}
	}
	return -1, -1
}

func (h Hand) Strait5() (int, int) {
	sort.Sort(h)
	for start := 0; start < len(h)-4; start++ {
		if h[start].suit == h[start+1].suit &&
			h[start].suit == h[start+2].suit &&
			h[start].suit == h[start+3].suit &&
			h[start].suit == h[start+4].suit &&
			h[start].rank+1 == h[start+1].rank &&
			h[start].rank+2 == h[start+2].rank &&
			h[start].rank+3 == h[start+3].rank &&
			h[start].rank+4 == h[start+4].rank {
			return h[start].suit, h[start].rank
		}
	}
	return -1, -1
}

func (a Hand) Pair() (int, int) {
	sort.Sort(a)
	for start := 0; start < len(a)-4; start++ {
		if a[start].suit == a[start+1].suit &&
			a[start].rank == a[start+1].rank {
			return a[start].suit, a[start].rank
		}
	}
	return -1, -1
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
