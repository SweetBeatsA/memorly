package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type User struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserInput struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserInput struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

type UserResponse struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var users = []User{
	{Id: 1, Name: "First User", Email: "first@gmail.com", Password: "firstPassword"},
	{Id: 2, Name: "Second User", Email: "second@gmail.com", Password: "secondPassword"},
	{Id: 3, Name: "Third User", Email: "third@gmail.com", Password: "thirdPassword"},
}

func createUser(c *gin.Context) {
	var input CreateUserInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user := User{Id: len(users) + 1, Name: input.Name, Email: input.Email, Password: input.Password}
	users = append(users, user)
	userResponse := user

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   userResponse,
	})
}

func getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  404,
			"message": "Not a proper id",
		})
		return
	}

	for _, user := range users {
		if user.Id == id {
			c.JSON(http.StatusOK, gin.H{
				"status": 200,
				"data":   user,
			})
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"status":  404,
		"message": "User not found",
	})
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   users,
	})
}

func updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	var input UpdateUserInput

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "No id",
		})
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	user := User{Id: id, Name: input.Name, Email: input.Email, Password: "password"}
	userResponse := user

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   userResponse,
	})
}

func deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "No id",
		})
		return
	}

	user := User{Id: id, Name: "deletedUser", Email: "deleted@google.com", Password: "password"}
	userResponse := user

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"data":   userResponse,
	})
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/users/:id", getUser)
	r.GET("/users", getUsers)
	r.POST("/user", createUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	return r
}

func main() {
	r := setupRouter()

	r.Run("localhost:8080")
}
