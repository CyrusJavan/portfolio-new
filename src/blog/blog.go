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

const (
	index           = "index.tpl"
	htmlContentType = "text/html; charset=utf-8"
	adminCookie     = "ADMIN_KEY"
)

var (
	notFoundBytes = []byte("<h1>404 Page not found :(</h1>")
)

type templateData struct {
	Year     int
	Page     string
	Title    string
	Articles []Article
	Article  Article
}

// Run defines the routes and starts the server.
func Run() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found: %w", err)
	}

	r := gin.Default()
	r.Static("/static", "static")

	t := template.Must(template.ParseGlob("tpl/*.tpl"))
	td := templateData{time.Now().UTC().Year(), "", "", nil, Article{}}

	r.GET("/", getRootHandler(t, td))
	r.GET("/about", getAboutHandler(t, td))
	r.GET("/blog", getBlogHandler(t, td))
	r.GET("/blog/:slug", getBlogSlugHandler(t, td))
	r.GET("/talks", getTalksHandler(t, td))

	authorized := r.Group("/")
	authorized.Use(getAuthMiddleware())
	authorized.GET("/edit/:slug", getEditSlugHandler(t, td))
	authorized.POST("/edit/:slug", getEditSlugPostHandler(t, td))

	r.POST("/api/track", handleAPITrack)

	r.NoRoute(func(c *gin.Context) {
		c.Data(http.StatusNotFound, htmlContentType, notFoundBytes)
	})

	p := os.Getenv("PORT")
	r.Run(":" + p)
}

func getRootHandler(t *template.Template, td templateData) func(*gin.Context) {
	return func(c *gin.Context) {
		td.Page = "Home"
		td.Title = "Software Engineer"
		td.Articles = getAllArticles()
		t.ExecuteTemplate(c.Writer, index, td)
	}
}

func getAboutHandler(t *template.Template, td templateData) func(*gin.Context) {
	return func(c *gin.Context) {
		td.Page = "About"
		td.Title = "About"
		t.ExecuteTemplate(c.Writer, index, td)
	}
}

func getBlogHandler(t *template.Template, td templateData) func(*gin.Context) {
	return func(c *gin.Context) {
		td.Page = "Blog"
		td.Title = "Blog"
		td.Articles = getAllArticles()
		t.ExecuteTemplate(c.Writer, index, td)
	}
}

func getBlogSlugHandler(t *template.Template, td templateData) func(*gin.Context) {
	return func(c *gin.Context) {
		s := c.Param("slug")
		a, err := getArticleBySlug(s)
		if err != nil {
			c.Data(http.StatusNotFound, htmlContentType, notFoundBytes)
			log.Printf("Slug not found: %s", s)
			return
		}
		td.Article = a
		td.Page = "BlogArticle"
		td.Title = a.Title
		t.ExecuteTemplate(c.Writer, index, td)
	}
}

func getTalksHandler(t *template.Template, td templateData) func(*gin.Context) {
	return func(c *gin.Context) {
		td.Page = "Talks"
		td.Title = "Talks"
		t.ExecuteTemplate(c.Writer, index, td)
	}
}

func getAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		k, err := c.Cookie(adminCookie)
		if err != nil || k != os.Getenv("ADMIN_KEY") {
			c.Data(http.StatusNotFound, htmlContentType, notFoundBytes)
			c.Abort()
			return
		}
		c.Next()
	}
}

func getEditSlugHandler(t *template.Template, td templateData) func(*gin.Context) {
	return func(c *gin.Context) {
		s := c.Param("slug")
		a, err := getArticleBySlug(s)
		if err != nil {
			c.Data(http.StatusNotFound, htmlContentType, notFoundBytes)
			log.Printf("Could not find article by slug to edit: %s", s)
			return
		}

		td.Article = a
		td.Page = "Editor"
		td.Title = "Editor"
		t.ExecuteTemplate(c.Writer, index, td)
	}
}

func getEditSlugPostHandler(t *template.Template, td templateData) func(*gin.Context) {
	return func(c *gin.Context) {
		s := c.Param("slug")
		err := updateArticleContent(s, c.PostForm("content"), c.PostForm("snippet"), c.PostForm("title"))
		if err != nil {
			log.Printf("Failed to edit article %s : %v", s, err)
			c.String(http.StatusBadRequest, "")
			return
		}
		c.String(http.StatusOK, "OK")
	}
}
