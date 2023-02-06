package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateInput struct {
	Id    int    `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required"`
}

func createUser(c *gin.Context) {
	input := CreateInput{}

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user := models.User{id: input.id, name: input.name, email: input.email, password: input.password}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   user,
	})
}

func getUser(c *gin.Context) {
	var id = c.Param("id")

	user := models.User{id: id, name: "test", email: "test@google.com"}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   user,
	})
}
