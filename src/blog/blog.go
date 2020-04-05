package blog

import (
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

// Run the server
func Run() {
	// load values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	// Heroku supplies the PORT env variable in production
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

	tmplData := struct {
		Year int
		Page string
	}{time.Now().UTC().Year(), ""}

	router.GET("/", func(c *gin.Context) {
		tmplData.Page = "Home"
		tmpl.ExecuteTemplate(c.Writer, "index.tpl", tmplData)
	})

	router.GET("/about", func(c *gin.Context) {
		tmplData.Page = "About"
		tmpl.ExecuteTemplate(c.Writer, "index.tpl", tmplData)
	})

	router.GET("/blog", func(c *gin.Context) {
		tmplData.Page = "Blog"
		tmpl.ExecuteTemplate(c.Writer, "index.tpl", tmplData)
	})

	router.GET("/talks", func(c *gin.Context) {
		tmplData.Page = "Talks"
		tmpl.ExecuteTemplate(c.Writer, "index.tpl", tmplData)
	})

	router.POST("/api/track", handleAPITrack)

	router.NoRoute(func(c *gin.Context) {
		c.Data(http.StatusNotFound, "text/html; charset=utf-8", []byte("<h1>404 Page not found :(</h1>"))
	})

	router.Run(":" + port)
}
