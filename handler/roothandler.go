package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RootHandler struct {
	ArticleHandler *ArticleHandler
}

func NewRootHandler(ArticleHandler *ArticleHandler) *RootHandler {
	return &RootHandler{ArticleHandler: ArticleHandler}
}

func (handler *RootHandler) ShowIndexPage(c *gin.Context) {
	articles := handler.ArticleHandler.GetAllArticle()

	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title":   "Home Page",
		"payload": articles}, "index.html")
}

func render(c *gin.Context, data gin.H, template string) {
	loggedInInterface, _ := c.Get("is_logged_in")
	data["is_logged_in"] = loggedInInterface.(bool)

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, template, data)
	}
}
