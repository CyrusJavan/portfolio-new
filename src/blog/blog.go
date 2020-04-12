package blog

import (
	"log"
	"net/http"
	"os"
	"text/template"
	"time"

	"github.com/CyrusJavan/portfolio-new/src/db"

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
		Year     int
		Page     string
		Title    string
		Articles []Article
		Article  Article
	}{time.Now().UTC().Year(), "", "", nil, Article{}}

	router.GET("/", func(c *gin.Context) {
		tmplData.Page = "Home"
		tmplData.Title = "Software Engineer"
		tmplData.Articles = getAllArticles()
		tmpl.ExecuteTemplate(c.Writer, "index.tpl", tmplData)
	})

	router.GET("/about", func(c *gin.Context) {
		tmplData.Page = "About"
		tmplData.Title = "About"
		tmpl.ExecuteTemplate(c.Writer, "index.tpl", tmplData)
	})

	router.GET("/blog", func(c *gin.Context) {
		tmplData.Page = "Blog"
		tmplData.Title = "Blog"
		tmplData.Articles = getAllArticles()
		tmpl.ExecuteTemplate(c.Writer, "index.tpl", tmplData)
	})

	router.GET("/blog/:slug", func(c *gin.Context) {
		slug := c.Param("slug")
		article := getArticleBySlug(slug)
		tmplData.Article = article
		tmplData.Page = "BlogArticle"
		tmplData.Title = article.Title
		tmpl.ExecuteTemplate(c.Writer, "index.tpl", tmplData)
	})

	router.GET("/talks", func(c *gin.Context) {
		tmplData.Page = "Talks"
		tmplData.Title = "Talks"
		tmpl.ExecuteTemplate(c.Writer, "index.tpl", tmplData)
	})

	htmlContentType := "text/html; charset=utf-8"
	notFoundBytes := []byte("<h1>404 Page not found :(</h1>")
	const adminCookie string = "ADMIN_KEY"

	router.GET("/editor/:slug", func(c *gin.Context) {
		key, err := c.Cookie(adminCookie)
		if err != nil {
			c.Data(http.StatusNotFound, htmlContentType, notFoundBytes)
			return
		}
		if key != os.Getenv("ADMIN_KEY") {
			c.Data(http.StatusNotFound, htmlContentType, notFoundBytes)
			return
		}

		slug := c.Param("slug")
		article := getArticleBySlug(slug)
		tmplData.Article = article
		tmplData.Page = "Editor"
		tmplData.Title = "Editor"
		tmpl.ExecuteTemplate(c.Writer, "index.tpl", tmplData)
	})

	router.POST("/editor/:slug", func(c *gin.Context) {
		key, err := c.Cookie(adminCookie)
		if err != nil {
			c.Data(http.StatusNotFound, htmlContentType, notFoundBytes)
			return
		}
		if key != os.Getenv("ADMIN_KEY") {
			c.Data(http.StatusNotFound, htmlContentType, notFoundBytes)
			return
		}

		slug := c.Param("slug")
		content := c.PostForm("content")
		snippet := c.PostForm("snippet")
		title := c.PostForm("title")
		err = updateArticleContent(slug, content, snippet, title)
		if err != nil {
			log.Printf("Failed to edit article %s : %v", slug, err)
			c.String(http.StatusBadRequest, "")
		}
		c.String(http.StatusOK, "OK")
	})

	router.POST("/api/track", handleAPITrack)

	router.NoRoute(func(c *gin.Context) {
		c.Data(http.StatusNotFound, htmlContentType, notFoundBytes)
	})

	router.Run(":" + port)
}

// Article information
type Article struct {
	ID          int     `db:"id"`
	Slug        string  `db:"slug"`
	Author      string  `db:"name"`
	Timestamp   float64 `db:"date_part"`
	Date        string
	AuthorImage string `db:"image_url"`
	Title       string `db:"title"`
	Content     string `db:"content"`
	Snippet     string `db:"snippet"`
	Tags        []string
}

func getAllArticles() []Article {
	db := db.GetInstance()
	sql := `
	SELECT ar.id, ar.slug, au.name, extract(epoch from ar.created_at), au.image_url, ar.title, ar.content, ar.snippet
	FROM article ar JOIN author au ON au.id=ar.author_id
	WHERE ar.is_active;
	`
	rows, err := db.Queryx(sql)
	if err != nil {
		log.Printf("Failed to getAllArticles: %s", err.Error())
	}
	articles := []Article{}
	for rows.Next() {
		article := Article{}
		err = rows.StructScan(&article)
		if err != nil {
			log.Printf("Failed to get scan a row into article struct: %s", err.Error())
			continue
		}
		ts := article.Timestamp
		secs := int64(ts)
		article.Date = time.Unix(secs, 0).Format("Monday, January 2 2006")
		article.Tags = getArticleTags(article.ID)
		articles = append(articles, article)
	}
	rows.Close()

	return articles
}

func getArticleTags(articleID int) []string {
	db := db.GetInstance()
	sql := `
	SELECT tag.name
	FROM article_tag JOIN tag ON tag.id = article_tag.tag_id
	WHERE article_tag.article_id = $1
	`
	rows, err := db.Queryx(sql, articleID)
	if err != nil {
		log.Printf("Failed to getArticleTags for articleID %d: %s", articleID, err.Error())
		return []string{}
	}
	tags := []string{}
	for rows.Next() {
		tag := struct {
			Name string
		}{}
		err = rows.StructScan(&tag)
		if err != nil {
			log.Printf("Failed to get scan a row into tag struct: %s", err.Error())
			continue
		}
		tags = append(tags, tag.Name)
	}
	rows.Close()

	return tags
}

func getArticleBySlug(slug string) Article {
	db := db.GetInstance()
	sql := `
	SELECT ar.id, ar.slug, au.name, extract(epoch from ar.created_at), au.image_url, ar.title, ar.content, ar.snippet
	FROM article ar JOIN author au ON au.id=ar.author_id
	WHERE ar.slug = $1;
	`
	article := Article{}
	err := db.Get(&article, sql, slug)
	if err != nil {
		log.Printf("Failed to getArticleBySlug slug %s: %s", slug, err.Error())
		return Article{}
	}

	ts := article.Timestamp
	secs := int64(ts)
	article.Date = time.Unix(secs, 0).Format("Monday, January 2 2006")
	article.Tags = getArticleTags(article.ID)

	return article
}

func updateArticleContent(slug, content, snippet, title string) error {
	db := db.GetInstance()
	sql := `
	UPDATE article
	SET content = :content, snippet = :snippet, title = :title
	WHERE slug = :slug
	`
	_, err := db.NamedExec(sql, map[string]interface{}{
		"content": content,
		"snippet": snippet,
		"slug":    slug,
		"title":   title,
	})
	return err
}
