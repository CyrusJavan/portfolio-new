package blog

import (
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

// Run the server
func Run() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

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
	tmpl := template.Must(template.ParseGlob("tpl/*.tpl"))

	router.GET("/", func(c *gin.Context) {
		tmpl.ExecuteTemplate(c.Writer, "index.tpl", nil)
	})

	router.POST("/api/track", handleAPITrack)

	router.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Sorry, this page is no where to be found. :(")
	})

	router.Run(":" + port)
}
