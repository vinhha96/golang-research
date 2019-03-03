package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vinhha96/golang-research/database"
	"github.com/vinhha96/golang-research/models"
	"github.com/vinhha96/golang-research/routes"
	"github.com/vinhha96/golang-research/utils"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.DebugMode)

	utils.InitializeConfiguration("", "env", "./config")

	port := viper.GetString("api.port")

	// Database:
	//db, err := database.GetDBConnection(
	//	viper.GetString("database.dialect"),
	//	viper.GetString("database.url"))
	db, err := database.GetDBConnection("mysql", "arun:password@tcp(172.17.0.1:3306)/golang_research?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Connect DB error")
	}

	// Create DB
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Article{})

	// Router:
	router = gin.Default()
	router.LoadHTMLGlob("./templates/*")
	routes.InitRoutes(router, db)

	_ = router.Run(":" + port)
}
