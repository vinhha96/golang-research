package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vinhha96/golang-research/routes"
)

var router *gin.Engine

func main() {
	// Set Gin to production mode
	gin.SetMode(gin.ReleaseMode)

	router = gin.Default()

	router.LoadHTMLGlob("../templates/*")

	routes.InitRoutes(router)

	_ = router.Run(":3000")
}
