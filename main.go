package main

import (
	"net/http"

	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := "8080"
	router := gin.Default()
	router.Static("/static", "static")

	router.Use(static.Serve("/", static.LocalFile("./static", true)))

	// Handle 404s
	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Sorry, this page is no where to be found. :(")
	})

	router.Run(":" + port)
}
