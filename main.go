package main

import (
	"log"
	"net/http"
	"os"

	"html/template"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	port := ""
	if os.Getenv("PORT") == "" {
		port = "7070"
	} else {
		port = os.Getenv("PORT")
	}

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	router.Static("/static", "static")
	tmpl := template.Must(template.ParseGlob("src/tpl/*.tpl"))

	router.GET("/", func(c *gin.Context) {
		tmpl.ExecuteTemplate(c.Writer, "index.tpl", nil)
	})

	//router.Use(static.Serve("/", static.LocalFile("./static", true)))

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Sorry, this page is no where to be found. :(")
	})

	router.Run(":" + port)
}
