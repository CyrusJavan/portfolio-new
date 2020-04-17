package blog

import (
	"time"

	"github.com/apex/log"
	"github.com/gin-gonic/gin"

	"github.com/CyrusJavan/portfolio-new/src/db"
)

// Article contains the information needed to display an article on the site.
type Article struct {
	ID          int     `db:"id"`
	Slug        string  `db:"slug"`
	Author      string  `db:"name"`
	AuthorImage string  `db:"image_url"`
	Title       string  `db:"title"`
	Content     string  `db:"content"`
	Snippet     string  `db:"snippet"`
	Timestamp   float64 `db:"date_part"`
	Date        string
	Tags        []string
}

func getAllArticles(c *gin.Context) []Article {
	db := db.GetInstance(c)
	sql := `
	SELECT ar.id, ar.slug, au.name, extract(epoch from ar.created_at), au.image_url, ar.title, ar.content, ar.snippet
	FROM article ar JOIN author au ON au.id=ar.author_id
	WHERE ar.is_active;
	`
	rows, err := db.Queryx(sql)
	if err != nil {
		l := c.Keys["logEntry"].(*log.Entry)
		l.WithError(err).WithFields(log.Fields{
			"func":  "getAllArticles",
			"query": sql,
		}).Error("query failed")
	}
	articles := []Article{}
	for rows.Next() {
		article := Article{}
		err = rows.StructScan(&article)
		if err != nil {
			l := c.Keys["logEntry"].(*log.Entry)
			l.WithError(err).WithFields(log.Fields{
				"func": "getAllArticles",
			}).Error("could not parse article query result")
			continue
		}
		ts := article.Timestamp
		secs := int64(ts)
		article.Date = time.Unix(secs, 0).Format("Monday, January 2 2006")
		article.Tags = getArticleTags(c, article.ID)
		articles = append(articles, article)
	}
	rows.Close()

	return articles
}

func getArticleTags(c *gin.Context, articleID int) []string {
	db := db.GetInstance(c)
	sql := `
	SELECT tag.name
	FROM article_tag JOIN tag ON tag.id = article_tag.tag_id
	WHERE article_tag.article_id = $1
	`
	rows, err := db.Queryx(sql, articleID)
	if err != nil {
		l := c.Keys["logEntry"].(*log.Entry)
		l.WithError(err).WithFields(log.Fields{
			"func":      "getArticleTags",
			"articleID": articleID,
		}).Error("article not found")
		return []string{}
	}
	tags := []string{}
	for rows.Next() {
		tag := struct {
			Name string
		}{}
		err = rows.StructScan(&tag)
		if err != nil {
			l := c.Keys["logEntry"].(*log.Entry)
			l.WithError(err).WithFields(log.Fields{
				"func":      "getArticleTags",
				"articleID": articleID,
			}).Error("could not parse article_tags query result")
			continue
		}
		tags = append(tags, tag.Name)
	}
	rows.Close()

	return tags
}

func getArticleBySlug(c *gin.Context, slug string) (Article, error) {
	db := db.GetInstance(c)
	sql := `
	SELECT ar.id, ar.slug, au.name, extract(epoch from ar.created_at), au.image_url, ar.title, ar.content, ar.snippet
	FROM article ar JOIN author au ON au.id=ar.author_id
	WHERE ar.slug = $1 AND ar.is_active;
	`
	article := Article{}
	err := db.Get(&article, sql, slug)
	if err != nil {
		l := c.Keys["logEntry"].(*log.Entry)
		l.WithError(err).WithFields(log.Fields{
			"func": "getArticleBySlug",
			"slug": slug,
		}).Error("article not found")
		return Article{}, err
	}

	ts := article.Timestamp
	secs := int64(ts)
	article.Date = time.Unix(secs, 0).Format("Monday, January 2 2006")
	article.Tags = getArticleTags(c, article.ID)

	return article, nil
}

func updateArticleContent(c *gin.Context, slug, content, snippet, title string) error {
	db := db.GetInstance(c)
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
