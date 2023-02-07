package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUserForm struct {
	Email    string `json:"email" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUserForm struct {
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

var users = []User{
	{Id: 1, Name: "First User", Email: "first@gmail.com", Password: "firstPassword"},
	{Id: 2, Name: "Second User", Email: "second@gmail.com", Password: "secondPassword"},
	{Id: 3, Name: "Third User", Email: "third@gmail.com", Password: "thirdPassword"},
}

func createUser(c *gin.Context) {
	var form CreateUserForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	highest := 0

	for _, user := range users {
		if user.Email == form.Email {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Email already used",
			})
			return
		}

		if user.Id > highest {
			highest = user.Id
		}
	}

	user := User{Id: highest + 1, Name: form.Name, Email: form.Email, Password: form.Password}
	users = append(users, user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully Created",
		"data":    user,
	})
}

func getUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Not a proper id",
		})
		return
	}

	for _, user := range users {
		if user.Id == id {
			c.JSON(http.StatusOK, gin.H{
				"data": user,
			})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "User not found",
	})
}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Not a proper id",
		})
		return
	}

	var form UpdateUserForm

	if err := c.ShouldBindJSON(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for _, user := range users {
		if user.Email == form.Email {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Email already used",
			})
			return
		}
	}

	for i, user := range users {
		if user.Id == id {
			users[i].Name = form.Name
			users[i].Email = form.Email

			c.JSON(http.StatusOK, gin.H{
				"message": "Successfully Updated",
				"data":    users[i],
			})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "User not found",
	})
}

func deleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"message": "Not a proper id",
		})
		return
	}

	for i, user := range users {
		if user.Id == id {
			users = append(users[:i], users[i+1:]...)

			c.JSON(http.StatusOK, gin.H{
				"message": "Successfully Deleted",
				"data":    user,
			})
			return
		}
	}

	c.JSON(http.StatusBadRequest, gin.H{
		"message": "User not found",
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
