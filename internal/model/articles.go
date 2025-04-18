package model

import (
	"database/sql"
	"html/template"
	"time"

	"github.com/elltja/news-website/internal/database"
)

type Article struct {
	ID        string        `json:"id"`
	Title     string        `json:"title"`
	Content   template.HTML `json:"content"`
	CreatedAt time.Time     `json:"created_at"`
}

type ArticleData struct {
	Title   string        `json:"title"`
	Content template.HTML `json:"content"`
}

func (a Article) FormattedCreatedAt() string {
	return a.CreatedAt.Format("Jan 2, 2006")
}

func GetArticles() ([]Article, error) {
	rows, err := database.DB.Query(`
		SELECT id, content, created_at, title FROM articles
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	articles := []Article{}
	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.ID, &article.Content, &article.CreatedAt, &article.Title); err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return articles, nil
}

func GetArticleById(id string) (Article, error) {
	row := database.DB.QueryRow(`
		SELECT id, title, content, created_at 
		FROM articles 
		WHERE id = $1`, id)

	var article Article

	if err := row.Scan(&article.ID, &article.Title, &article.Content, &article.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return article, nil
		}
		return article, err
	}
	return article, nil
}

func CreateArticle(data ArticleData) error {
	_, err := database.DB.Exec(`INSERT INTO articles (title, content) VALUES ($1, $2)`, data.Title, data.Content)
	if err != nil {
		return err
	}
	return nil
}

func DeleteArticle(id string) error {
	_, err := database.DB.Exec(`DELETE FROM articles WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
