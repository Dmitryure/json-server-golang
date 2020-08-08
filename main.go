package main

import (
	"fmt"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

var deck = CreateDeck()

func photoUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	c.SaveUploadedFile(file, path.Join("./photos", file.Filename))
	c.String(200, fmt.Sprintf("uploaded"))
}

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

// func getRandomNCards(c *gin.Context) {
// 	ShuffleDeck(deck)
// 	n, err := strconv.Atoi(c.Param("n"))
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	var hand []string
// 	if len(deck) <= n*2 {
// 		n = len(deck) - 1
// 	} else {
// 		fmt.Println(deck)
// 		for i := 0; i < n; i++ {
// 			fmt.Println(i)
// 			fmt.Println(hand, "hand")
// 			hand = append(hand, deck[i])
// 			deck = append(deck[:1], deck[2:]...)
// 		}
// 	}
// 	fmt.Println(len(deck))

// 	var typedHand = DeckStructs(hand)
// 	c.JSON(200, typedHand)
// }

func main() {
	r := gin.Default()
	r.POST("/photo", photoUpload)
	r.GET("/deck", getDeckString)
	r.GET("/jsondeck", getDeckJSON)
	// r.GET("/hand/:n", getRandomNCards)
	r.Run("0.0.0.0:3001") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
