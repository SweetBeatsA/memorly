package main

import (
	"memorly/configs"
	"memorly/routes"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())

	configs.ConnectDB()

	routes.AuthRoute(router)
	routes.UserRoute(router)

	return router
}

func main() {
	r := setupRouter()

	r.Run("localhost:8080")
}
