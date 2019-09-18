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
