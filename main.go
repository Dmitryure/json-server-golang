package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

var deck = CreateDeck()

func getDeckString(c *gin.Context) {
	ShuffleDeck(deck)
	c.String(200, strings.Join(deck, ", "))
}

func getDeckJSON(c *gin.Context) {
	ShuffleDeck(deck)
	var typedDeck = DeckStructs(deck)
	fmt.Println(len(typedDeck))
	c.JSON(200, typedDeck)
}

func getRandomNCards(c *gin.Context) {
	ShuffleDeck(deck)
	n, err := strconv.Atoi(c.Param("n"))
	if err != nil {
		fmt.Println(err)
	}
	var hand []string
	if len(deck) <= n*2 {
		n = len(deck) - 1
	} else {
		fmt.Println(deck)
		for i := 0; i < n; i++ {
			hand = append(hand, deck[i])
			deck = append(deck[:1], deck[2:]...)
		}
	}
	fmt.Println(len(deck))

	var typedHand = DeckStructs(hand)
	c.JSON(200, typedHand)
}

func generateNCards(n int) chan string {
	ch := make(chan string)

	go func(ch chan string) {
		for i := 0; i < n; i++ {
			if len(deck) == 0 {
				ch <- ""
			} else {
				ch <- deck[0]
				deck = append(deck[:0], deck[1:]...)
			}
		}
		close(ch)
	}(ch)
	return ch
}

func getNCards(c *gin.Context) {
	n, err := strconv.Atoi(c.Param("n"))
	if err != nil {
		fmt.Println(err)
	}
	var hand []string
	for val := range generateNCards(n) {
		hand = append(hand, val)
	}
	c.JSON(200, DeckStructs(hand))
}

func resetDeck(c *gin.Context) {
	deck = CreateDeck()
}

func main() {
	r := gin.Default()
	fmt.Println("server is running on 8080")
	r.GET("/deck", getDeckString)
	r.GET("/jsondeck", getDeckJSON)
	r.GET("/hand/:n", getRandomNCards)
	r.GET("/gdeck/:n", getNCards)
	r.GET("/reset", resetDeck)
	r.Run("0.0.0.0:8080")
}
