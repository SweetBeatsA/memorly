package main

import (
	"github.com/gin-gonic/gin"
	"memorly/configs"
	"memorly/routes"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	configs.ConnectDB()

	routes.UserRoute(r)

	return r
}

func main() {
	r := setupRouter()

	r.Run("localhost:8080")
}
