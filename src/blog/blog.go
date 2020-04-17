package blog

import (
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/text"
	"github.com/google/uuid"

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
	log.SetHandler(text.New(os.Stderr))

	if err := godotenv.Load(); err != nil {
		log.WithError(err).WithField("func", "Run").Info("No .env file found")
	}

	r := gin.New()
	r.Use(loggerMiddleware())
	r.Use(gin.Recovery())
	r.Static("/static", "static")
	authorized := r.Group("/")
	authorized.Use(authMiddleware())

	t := template.Must(template.ParseGlob("tpl/*.tpl"))
	td := templateData{time.Now().UTC().Year(), "", "", nil, Article{}}

	r.GET("/", rootHandler(t, td))
	r.GET("/about", aboutHandler(t, td))
	r.GET("/blog", blogHandler(t, td))
	r.GET("/blog/:slug", blogSlugHandler(t, td))
	r.GET("/talks", talksHandler(t, td))
	authorized.GET("/edit/:slug", editSlugHandler(t, td))
	authorized.POST("/edit/:slug", editSlugPostHandler(t, td))

	r.POST("/api/track", handleAPITrack)

	r.NoRoute(func(c *gin.Context) {
		c.Data(http.StatusNotFound, htmlContentType, notFoundBytes)
	})

	p := os.Getenv("PORT")
	r.Run(":" + p)
}

func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		if c.Keys == nil {
			c.Keys = make(map[string]interface{})
		}
		c.Keys["logEntry"] = log.WithField("reqID", uuid.New())
		c.Next()

		end := time.Now()
		latency := end.Sub(start)
		entry := c.Keys["logEntry"].(*log.Entry)
		entry = entry.WithFields(log.Fields{
			"status":     c.Writer.Status(),
			"method":     c.Request.Method,
			"path":       c.Request.URL.Path,
			"ip":         c.ClientIP(),
			"latency":    latency,
			"user-agent": c.Request.UserAgent(),
			"time":       end.Format(time.UnixDate),
		})

		if len(c.Errors) > 0 {
			entry.Error(c.Errors.String())
		} else {
			entry.Info("")
		}
	}
}

func rootHandler(t *template.Template, td templateData) func(*gin.Context) {
	return func(c *gin.Context) {
		td.Page = "Home"
		td.Title = "Software Engineer"
		td.Articles = getAllArticles(c)
		err := t.ExecuteTemplate(c.Writer, index, td)
		if err != nil {
			l := c.Keys["logEntry"].(*log.Entry)
			l.WithError(err).WithFields(log.Fields{
				"func":         "rootHandler",
				"template":     index,
				"templateData": td,
			}).Error("template failed to execute")
		}
	}
}

func aboutHandler(t *template.Template, td templateData) func(*gin.Context) {
	return func(c *gin.Context) {
		td.Page = "About"
		td.Title = "About"
		err := t.ExecuteTemplate(c.Writer, index, td)
		if err != nil {
			l := c.Keys["logEntry"].(*log.Entry)
			l.WithError(err).WithFields(log.Fields{
				"func":         "aboutHandler",
				"template":     index,
				"templateData": td,
			}).Error("template failed to execute")
		}
	}
}

func blogHandler(t *template.Template, td templateData) func(*gin.Context) {
	return func(c *gin.Context) {
		td.Page = "Blog"
		td.Title = "Blog"
		td.Articles = getAllArticles(c)
		err := t.ExecuteTemplate(c.Writer, index, td)
		if err != nil {
			l := c.Keys["logEntry"].(*log.Entry)
			l.WithError(err).WithFields(log.Fields{
				"func":         "blogHandler",
				"template":     index,
				"templateData": td,
			}).Error("template failed to execute")
		}
	}
}

func blogSlugHandler(t *template.Template, td templateData) func(*gin.Context) {
	return func(c *gin.Context) {
		s := c.Param("slug")
		a, err := getArticleBySlug(c, s)
		if err != nil {
			c.Data(http.StatusNotFound, htmlContentType, notFoundBytes)
			l := c.Keys["logEntry"].(*log.Entry)
			l.WithError(err).WithFields(log.Fields{
				"func": "blogSlugHandler",
				"slug": s,
			}).Error("article not found")
			return
		}
		td.Article = a
		td.Page = "BlogArticle"
		td.Title = a.Title
		err = t.ExecuteTemplate(c.Writer, index, td)
		if err != nil {
			l := c.Keys["logEntry"].(*log.Entry)
			l.WithError(err).WithFields(log.Fields{
				"func":         "blogSlugHandler",
				"template":     index,
				"templateData": td,
			}).Error("template failed to execute")
		}
	}
}

func talksHandler(t *template.Template, td templateData) func(*gin.Context) {
	return func(c *gin.Context) {
		td.Page = "Talks"
		td.Title = "Talks"
		err := t.ExecuteTemplate(c.Writer, index, td)
		if err != nil {
			l := c.Keys["logEntry"].(*log.Entry)
			l.WithError(err).WithFields(log.Fields{
				"func":         "talksHandler",
				"template":     index,
				"templateData": td,
			}).Error("template failed to execute")
		}
	}
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		k, err := c.Cookie(adminCookie)
		if err != nil || k != os.Getenv("ADMIN_KEY") {
			c.Data(http.StatusNotFound, htmlContentType, notFoundBytes)
			l := c.Keys["logEntry"].(*log.Entry)
			l.WithError(err).WithFields(log.Fields{
				"func": "authMiddleware",
			}).Error("auth failed")
			c.Abort()
			return
		}
		c.Next()
	}
}

func editSlugHandler(t *template.Template, td templateData) func(*gin.Context) {
	return func(c *gin.Context) {
		s := c.Param("slug")
		a, err := getArticleBySlug(c, s)
		if err != nil {
			c.Data(http.StatusNotFound, htmlContentType, notFoundBytes)
			l := c.Keys["logEntry"].(*log.Entry)
			l.WithError(err).WithFields(log.Fields{
				"func": "editSlugHandler",
				"slug": s,
			}).Error("article not found")
			return
		}

		td.Article = a
		td.Page = "Editor"
		td.Title = "Editor"
		err = t.ExecuteTemplate(c.Writer, index, td)
		if err != nil {
			l := c.Keys["logEntry"].(*log.Entry)
			l.WithError(err).WithFields(log.Fields{
				"func":         editSlugHandler,
				"template":     index,
				"templateData": td,
			}).Error("template failed to execute")
		}
	}
}

func editSlugPostHandler(t *template.Template, td templateData) func(*gin.Context) {
	return func(c *gin.Context) {
		s := c.Param("slug")
		err := updateArticleContent(c, s, c.PostForm("content"), c.PostForm("snippet"), c.PostForm("title"))
		if err != nil {
			l := c.Keys["logEntry"].(*log.Entry)
			l.WithError(err).WithFields(log.Fields{
				"func": "editSlugHandler",
				"slug": s,
			}).Error("failed to update article content")
			c.String(http.StatusBadRequest, "")
			return
		}
		c.String(http.StatusOK, "OK")
	}
}
