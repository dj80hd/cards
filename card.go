package main

var strRanks = []string{"2", "3", "4", "5", "6", "7", "8", "9", "T", "J", "Q", "K", "A"}
var strSuits = []string{"c", "d", "h", "s"}

// Card is a playing card with a suit and rank
type Card struct {
	suit int
	rank int
}
