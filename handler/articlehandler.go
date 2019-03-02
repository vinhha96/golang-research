package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/vinhha96/golang-research/models"
	"net/http"
	"strconv"
)

type ArticleHandler struct {
	db *gorm.DB
}

func NewArticleHandler(db *gorm.DB) *ArticleHandler {
	return &ArticleHandler{db: db}
}

func (articleHandler *ArticleHandler) GetArticleByID(ctx *gin.Context) {
	if articleID, err := strconv.Atoi(ctx.Param("article_id")); err == nil {
		if article, err := articleHandler.GetArticleByIDFromDB(articleID); err == nil {
			render(ctx, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")
		} else {
			_ = ctx.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}

func (articleHandler *ArticleHandler) ShowCreateArticlePage(ctx *gin.Context) {
	render(ctx, gin.H{
		"title": "Create new article",
	}, "create-article.html")
}

func (articleHandler *ArticleHandler) CreateNewArticle(ctx *gin.Context) {
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")

	if article, err := articleHandler.CreateNewArticleInDB(title, content); err == nil {
		render(ctx, gin.H{
			"title":   "Submission successfuly",
			"payload": article,
		}, "submission-successful.html")
	} else {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
}

func (articleHandler *ArticleHandler) GetArticleByIDFromDB(articleID int) (*models.Article, error) {
	var articleList []models.Article

	articleHandler.db.Find(&articleList)

	for _, article := range articleList {
		if article.ID == articleID {
			return &article, nil
		}
	}
	return nil, errors.New("Can't find article in database")
}

func (articleHandler *ArticleHandler) CreateNewArticleInDB(title, content string) (*models.Article, error) {
	var articleList []models.Article
	articleHandler.db.Find(&articleList)

	newArticle := models.Article{ID: len(articleList) + 1, Title: title, Content: content}

	articleList = append(articleList, newArticle)

	return &newArticle, nil
}
