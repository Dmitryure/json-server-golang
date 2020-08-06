package main

import (
	"fmt"
	"path"

	"github.com/gin-gonic/gin"
)

func photoUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	c.SaveUploadedFile(file, path.Join("./photos", file.Filename))
	c.String(200, fmt.Sprintf("uploaded"))
}

func getDeckString(c *gin.Context) {
	c.String(200, DeckString)
}

func getDeckJSON(c *gin.Context) {
	c.JSON(200, TypedDeck)
}

func main() {
	r := gin.Default()
	r.POST("/photo", photoUpload)
	r.GET("/deck", getDeckString)
	r.GET("/jsondeck", getDeckJSON)
	r.Run("0.0.0.0:3001") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
