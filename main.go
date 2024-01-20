package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", nil)
	})

	fortunes := []string{
		"Good things come to those who wait, but better things come to those who work for it.",
		"An exciting opportunity is just around the corner.",
		"Don't be afraid to take risks – fortune favors the bold.",
		"A smile is the universal language of kindness.",
		"Your future is as bright as the stars.",
	}

	r.GET("/generate-fortune", func(ctx *gin.Context) {
		rand.Seed(time.Now().UnixNano())
		randomIndex := rand.Intn(len(fortunes))

		// get fortune
		fortune := fortunes[randomIndex]

		// return the fortune as JSON
		ctx.JSON(http.StatusOK, gin.H{"fortune": fortune})
	})

	r.Run(":8080")
}
