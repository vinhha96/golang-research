package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/vinhha96/golang-research/routes"
	"github.com/vinhha96/golang-research/utils"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.DebugMode)

	utils.InitializeConfiguration("", "env", "../config")

	port := viper.GetString("api.port")

	router = gin.Default()

	router.LoadHTMLGlob("../templates/*")

	routes.InitRoutes(router)

	_ = router.Run(":" + port)
}
