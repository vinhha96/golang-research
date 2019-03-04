package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/vinhha96/golang-research/handler"
	"github.com/vinhha96/golang-research/models"
	"github.com/vinhha96/golang-research/storages"
)

func InitRoutes(router *gin.Engine, db *gorm.DB, redisClient *storages.RedisClient) {

	router.Use(handler.SetUserStatus())
	router.Static("/views", "./views")

	userHandler := handler.NewUserHandler()
	userHandler.UserStore = models.NewUserStore(db)

	articleHandler := handler.NewArticleHandler()
	articleHandler.ArticleStore = models.NewArticleStore(db)

	rootHandler := handler.NewRootHandler(articleHandler)

	router.GET("/", rootHandler.ShowIndexPage)

	userRoutes := router.Group("/u")
	{
		userRoutes.Static("/views", "./views")

		userRoutes.GET("/user/:userID", userHandler.ShowProfilePage)

		userRoutes.GET("/login", handler.EnsureNotLoggedIn(), userHandler.LoginPage)

		userRoutes.POST("/login", handler.EnsureNotLoggedIn(), userHandler.PerformLogin)

		userRoutes.GET("/logout", userHandler.LogoutPage)

		userRoutes.GET("/register", userHandler.ShowRegistrationPage)

		userRoutes.POST("/register", userHandler.Register)
	}

	articleRoutes := router.Group("/article")
	{
		articleRoutes.Static("/views", "./views")

		articleRoutes.GET("/view/:article_id", articleHandler.GetArticleByID)

		articleRoutes.GET("/create", handler.EnsureLoggedIn(), articleHandler.ShowCreateArticlePage)

		articleRoutes.POST("/create", handler.EnsureLoggedIn(), articleHandler.CreateNewArticle)
	}
}
