package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/vinhha96/golang-research/models"
	"net/http"
	"strconv"
)

func GetArticleByID(ctx *gin.Context) {
	if articleID, err := strconv.Atoi(ctx.Param("article_id")); err == nil {
		if article, err := models.GetArticleByID(articleID); err == nil {
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

func ShowCreateArticlePage(ctx *gin.Context) {
	render(ctx, gin.H{
		"title": "Create new article",
	}, "create-article.html")
}

func CreateNewArticle(ctx *gin.Context) {
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")

	if article, err := models.CreateNewArticle(title, content); err == nil {
		render(ctx, gin.H{
			"title":   "Submission successfuly",
			"payload": article,
		}, "submission-successful.html")
	} else {
		ctx.AbortWithStatus(http.StatusBadRequest)
	}

}
