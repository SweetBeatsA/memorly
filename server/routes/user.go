package routes

import (
	"memorly/controllers"
	"memorly/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.Use(middleware.Authenticate())
	router.GET("/user", controllers.GetUser())
}
