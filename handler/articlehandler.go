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

}

func CreateNewArticle(ctx *gin.Context) {

}
