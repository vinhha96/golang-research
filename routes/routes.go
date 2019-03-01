package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vinhha96/golang-research/handler"
)

func InitRoutes(router *gin.Engine) {

	router.Use(handler.SetUserStatus())

	router.GET("/", handler.ShowIndexPage)

	userRoutes := router.Group("/u")
	{
		userRoutes.GET("/user/:userID", handler.ShowProfilePage)

		userRoutes.GET("/login", handler.EnsureNotLoggedIn(), handler.LoginPage)

		userRoutes.POST("/login", handler.EnsureNotLoggedIn(), handler.PerformLogin)

		userRoutes.GET("/logout", handler.LogoutPage)

		userRoutes.GET("/register", handler.ShowRegistrationPage)

		userRoutes.POST("/register", handler.Register)
	}

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", handler.GetArticleByID)

		articleRoutes.GET("/create", handler.EnsureLoggedIn(), handler.ShowCreateArticlePage)

		articleRoutes.POST("/create", handler.EnsureLoggedIn(), handler.CreateNewArticle)
	}
}
