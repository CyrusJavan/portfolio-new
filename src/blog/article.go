package blog

import (
	"log"
	"time"

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

func getArticleBySlug(slug string) (Article, error) {
	db := db.GetInstance()
	sql := `
	SELECT ar.id, ar.slug, au.name, extract(epoch from ar.created_at), au.image_url, ar.title, ar.content, ar.snippet
	FROM article ar JOIN author au ON au.id=ar.author_id
	WHERE ar.slug = $1 AND ar.is_active;
	`
	article := Article{}
	err := db.Get(&article, sql, slug)
	if err != nil {
		log.Printf("Failed to getArticleBySlug slug %s: %s", slug, err.Error())
		return Article{}, err
	}

	ts := article.Timestamp
	secs := int64(ts)
	article.Date = time.Unix(secs, 0).Format("Monday, January 2 2006")
	article.Tags = getArticleTags(article.ID)

	return article, nil
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
