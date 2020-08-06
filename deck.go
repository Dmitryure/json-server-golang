package main

import (
	"fmt"
	"strings"
)

var s = []string{"Spades", "Hearts", "Diamonds", "Clubs"}
var v = []string{"Ace", "Deuce", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

// CreateDeck creates deck
func CreateDeck() (d []string) {
	for _, suit := range s {
		for _, value := range v {
			d = append(d, value+" of "+suit)
		}
	}

	fmt.Println(d)
	return
}

var deck = CreateDeck()

// Card type
type Card struct {
	Info  string `json:"info"`
	Suit  string `json:"suit"`
	Value int    `json:"value"`
}

// Deck type

// JSONdeck returns typed deck
func JSONdeck(d []string) (jsond []Card) {
	var v int
	var s string
	for _, card := range d {
		// G E N I U S
		switch true {
		case strings.Contains(card, "Spades"):
			s = "spades"
		case strings.Contains(card, "Hearts"):
			s = "hearts"
		case strings.Contains(card, "Diamonds"):
			s = "diamonds"
		case strings.Contains(card, "Clubs"):
			s = "clubs"
		}
		switch true {
		case strings.Contains(card, "Ace"):
			v = 14
		case strings.Contains(card, "Deuce"):
			v = 2
		case strings.Contains(card, "Three"):
			v = 3
		case strings.Contains(card, "Four"):
			v = 4
		case strings.Contains(card, "Five"):
			v = 5
		case strings.Contains(card, "Six"):
			v = 6
		case strings.Contains(card, "Seven"):
			v = 7
		case strings.Contains(card, "Eight"):
			v = 8
		case strings.Contains(card, "Nine"):
			v = 9
		case strings.Contains(card, "Ten"):
			v = 10
		case strings.Contains(card, "Jack"):
			v = 11
		case strings.Contains(card, "Queen"):
			v = 12
		case strings.Contains(card, "King"):
			v = 13
		}
		jsond = append(jsond, Card{Info: card, Suit: s, Value: v})
	}
	return
}

// DeckString returns string representation of deck
var DeckString = strings.Join(deck, ", ")

// TypedDeck typed deck
var TypedDeck = JSONdeck(deck)

// GetNRandomCards get n random cards from deck
func GetNRandomCards(n int, d []string) {
	fmt.Println(len(d))
	for i := 0; i < n; i++ {

	}
}
