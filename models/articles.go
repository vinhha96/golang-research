package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Article struct {
	ID        int `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Title   string `json:"title"`
	Content string `json:"content" sql:"type:varchar(10000)"`
}

type ArticleStore struct {
	db *gorm.DB
}

func NewArticleStore(db *gorm.DB) *ArticleStore {
	return &ArticleStore{db: db}
}

func (s *ArticleStore) GetAll() []Article {
	return nil
}

func (s *ArticleStore) GetArticleByIDFromDB(articleID int) (*Article, error) {
	var articleList []Article
	s.db.Find(&articleList)

	for _, article := range articleList {
		if article.ID == articleID {
			return &article, nil
		}
	}
	return nil, errors.New("Can't find article in database")
}

func (s *ArticleStore) CreateNewArticleInDB(title, content string) (*Article, error) {
	var articleList []Article
	s.db.Find(&articleList)

	newArticle := Article{ID: len(articleList) + 1, Title: title, Content: content}

	s.db.Save(&newArticle)

	return &newArticle, nil
}

func (s *ArticleStore) GetAllArticle() *[]Article {
	var articleList []Article
	s.db.Find(&articleList)
	return &articleList
}
