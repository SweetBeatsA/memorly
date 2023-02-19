package main

import (
	"gin/configs"
	"gin/routes"
	"github.com/gin-gonic/gin"
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
