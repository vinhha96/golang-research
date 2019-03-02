package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vinhha96/golang-research/models"
	"net/http"
	"strconv"
)

type ArticleHandler struct {
	ArticleStore *models.ArticleStore
}

func NewArticleHandler() *ArticleHandler {
	return &ArticleHandler{}
}

func (articleHandler *ArticleHandler) GetAllArticle() *[]models.Article {
	return articleHandler.ArticleStore.GetAllArticle()
}

func (articleHandler *ArticleHandler) GetArticleByID(ctx *gin.Context) {
	if articleID, err := strconv.Atoi(ctx.Param("article_id")); err == nil {
		if article, err := articleHandler.ArticleStore.GetArticleByIDFromDB(articleID); err == nil {
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

	if article, err := articleHandler.ArticleStore.CreateNewArticleInDB(title, content); err == nil {
		render(ctx, gin.H{
			"title":   "Submission successfuly",
			"payload": article,
		}, "submission-successful.html")
	} else {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}
}
