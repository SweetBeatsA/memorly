package main

import (
	"gin/configs"
	"gin/controllers"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	configs.ConnectDB()

	/* USER */
	user := new(controllers.UserController)

	r.POST("user/login", user.login)
	r.POST("/user/register", user.register)
	r.GET("/users", user.getAll)
	r.PUT("/users/:id", user.update)
	r.DELETE("/users/:id", user.delete)

	return r
}

func main() {
	r := setupRouter()

	r.Run("localhost:8080")
}
