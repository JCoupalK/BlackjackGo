package main

import (
	"math/rand"
	"time"
)

type Card struct {
	Rank    string
	Suit    string
	Symbols string
}

// newDeck creates a new deck of 52 playing cards
func newDeck() []Card {
	suits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	ranks := []string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
	var deck []Card

	for _, suit := range suits {
		for _, rank := range ranks {
			deck = append(deck, Card{Rank: rank, Suit: suit})
		}
	}

	return deck
}

// shuffleDeck shuffles the given deck of cards
func shuffleDeck(deck []Card) []Card {
	rand.NewSource(time.Now().UnixNano())
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
	return deck
}

// dealCard deals a card from the deck and returns the card along with the remaining deck
func dealCard(deck []Card) (Card, []Card) {
	card := deck[0]       // Take the top card
	return card, deck[1:] // Return the card and the rest of the deck
}
