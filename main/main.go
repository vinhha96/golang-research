package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vinhha96/golang-research/database"
	"github.com/vinhha96/golang-research/models"
	"github.com/vinhha96/golang-research/routes"
	"github.com/vinhha96/golang-research/storages"
	"github.com/vinhha96/golang-research/utils"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.DebugMode)

	utils.InitializeConfiguration("", "env", "../config")

	port := viper.GetString("api.port")

	fmt.Println(fmt.Sprintf("Dialect: %s", viper.GetString("database.dialect")))
	fmt.Println(fmt.Sprintf("Url: %s", viper.GetString("database.url")))

	// Database:
	db, err := database.GetDBConnection(
		viper.GetString("database.dialect"),
		viper.GetString("database.url"))

	if err != nil {
		panic("Connect DB error")
	}

	// Create DB
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Article{})

	// Create Redis:

	redisClient := storages.GetRedisClient()
	redisClient.SaveToStore()

	// Router:
	router = gin.Default()
	router.LoadHTMLGlob("../templates/*")
	routes.InitRoutes(router, db)

	_ = router.Run(":" + port)
}
