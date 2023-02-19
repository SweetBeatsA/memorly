package routes

import (
	"github.com/gin-gonic/gin"
	"memorly/controllers"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user", controllers.CreateUser())
	router.GET("/users/:userId", controllers.GetAUser())
}
