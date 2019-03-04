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

	utils.InitializeConfiguration("", "env", "./config")

	fmt.Println(fmt.Sprintf("[Config] API Port: %s", viper.GetString("api.port")))
	port := viper.GetString("api.port")

	fmt.Println(fmt.Sprintf("[Config] Dialect: %s", viper.GetString("database.dialect")))
	fmt.Println(fmt.Sprintf("[Config] Url: %s", viper.GetString("database.url")))

	// Connect Database:
	db, err := database.GetDBConnection(
		viper.GetString("database.dialect"),
		viper.GetString("database.url"))

	if err != nil {
		panic("[Error] Connect DB error")
	}

	// Create DB
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Article{})

	// Create Redis:
	fmt.Println(fmt.Sprintf("[Config] Redis address: %s", viper.GetString("redis.address")))
	fmt.Println(fmt.Sprintf("[Config] Redis password: %s", viper.GetString("redis.password")))

	redisClient := storages.GetRedisClient(
		viper.GetString("redis.address"),
		viper.GetString("redis.password"))

	// Router:
	router = gin.Default()
	router.LoadHTMLGlob("./templates/*")
	routes.InitRoutes(router, db, redisClient)

	_ = router.Run(":" + port)
}
