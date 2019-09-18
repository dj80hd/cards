package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var strRanks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var strSuits = []string{"c", "d", "h", "s"}

type Hand []Card

func (a Hand) Len() int           { return len(a) }
func (a Hand) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Hand) Less(i, j int) bool { return a[i].suit < a[j].suit || a[i].rank < a[i].rank }

// Three returns the rank for any 3 of a kind, -1 if none/error
func (a Hand) Three() int {
	if a == nil {
		return -1
	}

	var trios = [10][3]int{
		{0, 1, 2},
		{0, 1, 3},
		{0, 1, 4},
		{0, 2, 3},
		{0, 2, 4},
		{0, 3, 4},
		{1, 2, 3},
		{1, 2, 4},
		{1, 3, 4},
		{2, 3, 4},
	}

	for t := 0; t < 10; t++ {
		if a[trios[t][0]].rank == a[trios[t][1]].rank &&
			a[trios[t][0]].rank == a[trios[t][2]].rank {
			return a[trios[t][0]].rank
		}
	}

	return -1
}

//Strait4 returns the suit and rank of a straight. -1 for none
func (a Hand) Strait4() (int, int) {
	for start := 0; start < len(a)-3; start++ {
		if a[start].suit == a[start+1].suit &&
			a[start].suit == a[start+2].suit &&
			a[start].suit == a[start+3].suit &&
			a[start].rank+1 == a[start+1].rank &&
			a[start].rank+2 == a[start+2].rank &&
			a[start].rank+3 == a[start+3].rank {
			return a[start].suit, a[start].rank
		}
	}
	return -1, -1
}

// NewCard replaces the nth card with a new one from the deck
func (h Hand) NewCard(n int, d *Deck) error {
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

// Card is a playing card with a suit and rank
type Card struct {
	suit int
	rank int
}

// Deck is a group of cards that can be drawn out
type Deck struct {
	cards []Card
	index int
	mux   sync.Mutex
}

func (d *Deck) Len() int {
	return len(d.cards) - d.index
}

func (d *Deck) Draw(n int) ([]Card, error) {
	d.mux.Lock()
	defer d.mux.Unlock()
	if n > d.Len() {
		return nil, fmt.Errorf("draw %d > len %d", n, d.Len())
	}

	cards := make([]Card, 0)
	for i := 0; i < n; i++ {
		d.index = d.index + 1
		cards = append(cards, d.cards[d.index-1])
	}
	return cards, nil
}

func (d *Deck) String() string {
	return fmt.Sprintf("%v at %d", d.cards, d.index)
}

func NewDeck() *Deck {
	c := make([]Card, 0)
	for s := 0; s < len(strSuits); s++ {
		for r := 0; r < len(strRanks); r++ {
			c = append(c, Card{suit: s, rank: r})
		}
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(c), func(i, j int) { c[i], c[j] = c[j], c[i] })
	return &Deck{cards: c}
}

func main() {
	fmt.Println("cards")
}
