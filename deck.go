package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
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

// Card type
type Card struct {
	Info  string `json:"info"`
	Suit  string `json:"suit"`
	Value int    `json:"value"`
}

// DeckStructs returns typed deck
func DeckStructs(d []string) (jsond []Card) {
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
		default:
			s = ""
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
		default:
			v = 0
		}
		jsond = append(jsond, Card{Info: card, Suit: s, Value: v})
	}
	return
}

// ShuffleDeck shuffles deck
func ShuffleDeck(d []string) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for n := len(d); n > 0; n-- {
		randIndex := r.Intn(n)
		d[n-1], d[randIndex] = d[randIndex], d[n-1]
	}
}
