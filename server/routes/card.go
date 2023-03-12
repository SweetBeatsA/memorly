package routes

import (
	"memorly/controllers"
	"memorly/middleware"

	"github.com/gin-gonic/gin"
)

func CardRoute(router *gin.Engine) {
	router.Use(middleware.Authenticate())
	router.POST("/card", controllers.CreateCard())
}