package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vinhha96/golang-research/handler"
)

func InitRoutes(router *gin.Engine) {
	router.GET("/", handler.ShowIndexPage)
}

