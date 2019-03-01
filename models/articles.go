package models

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// For this demo, we're storing the article list in memory
// In a real application, this list will most likely be fetched
// from a database or from static files
var articleList = []article{
	{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// Return a list of all the articles
func GetAllArticles() []article {
	return articleList
}

func GetArticleByID(articleID int) (*article, error) {
	for _, article := range articleList {
		if article.ID == articleID {
			return &article, nil
		}
	}
	return nil, errors.New("Can't find article in database")
}

func CreateNewArticle(title, content string) (*article, error) {
	newArticle := article{ID: len(articleList) + 1, Title: title, Content: content}

	articleList = append(articleList, newArticle)

	return &newArticle, nil
}
