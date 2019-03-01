package models

import (
	"errors"
	"time"
)

type Article struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the Article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []Article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func GetAllArticles() []Article {
	return articleList
}

func GetArticleByID(articleID int) (*Article, error) {
	for _, article := range articleList {
		if article.ID == articleID {
			return &article, nil
		}
	}
	return nil, errors.New("Can't find article in database")
}

func CreateNewArticle(title, content string) (*Article, error) {
	newArticle := Article{ID: len(articleList) + 1, Title: title, Content: content}

	articleList = append(articleList, newArticle)

	return &newArticle, nil
}
