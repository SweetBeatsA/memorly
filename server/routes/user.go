package routes

import (
	"memorly/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/users/:userId", controllers.GetAUser())
}
