package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := ""
	if os.Getenv("PORT") == "" {
		port = "8080"
	} else {
		port = os.Getenv("PORT")
	}

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	router.Static("/static", "static")

	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	// Handle 404s
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Sorry, this page is no where to be found. :(")
	})

	router.Run(":" + port)
}
