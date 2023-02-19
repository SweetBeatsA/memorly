package controllers

import (
	"gin/forms"
	"gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

var userModel = new(models.UserModel)

// var userForm = new(models.userForm)

// func login(c *gin.Context) {
// 	var form forms.LoginForm

// 	if err := c.ShouldBindJSON(&form); err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
// 			"message": "Input format is not valid",
// 		})
// 		return
// 	}

// 	token, err := userModel.login(form)

// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
// 			"message": "Input data is not matched",
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "Succesfully Log-in",
// 		"token":   token,
// 	})
// }

func register(c *gin.Context) {
	var form forms.RegisterForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "Input format is not valid",
		})
		return
	}

	user, err := userModel.register(form)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Registerd",
		"user":    user,
	})
}

func getAll(c *gin.Context) {
	users, err := userModel.getAll()

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
			"message": "No users to show",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Get All",
		"users":   users,
	})
}

// func update(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{
// 			"message": "Not a proper id",
// 		})
// 		return
// 	}

// 	var form forms.UpdateForm

// 	user, err := userModel.update

// 	if err := c.ShouldBindJSON(&form); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": err.Error(),
// 		})

// 		return
// 	}

// 	user := models.User{Id: id, Name: input.Name, Email: input.Email, Password: "password"}
// 	userResponse := user

// 	c.JSON(http.StatusOK, gin.H{
// 		"status": 200,
// 		"data":   userResponse,
// 	})
// }

// func delete(c *gin.Context) {
// 	id, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  400,
// 			"message": "No id",
// 		})
// 		return
// 	}

// 	user := models.User{Id: id, Name: "deletedUser", Email: "deleted@google.com", Password: "password"}
// 	userResponse := user

// 	c.JSON(http.StatusOK, gin.H{
// 		"status": 200,
// 		"data":   userResponse,
// 	})
// }
