package main

import (
    "bytes"
    "encoding/json"
    "io/ioutil"
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/rs/xid"
    "github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/user", func(c *gin.Context) {
		c.String(http.StatusOK, "successful")
	})

	return router
}

func TestGetUser(t *testing.T) {
	router := setupRouter()

	//request to /login
	request, err := http.NewRequest("GET", "/user", nil)
	assert.NoError(t, err)

	//record response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	//checking if response 200 OK is ok
	assert.Equal(t, http.StatusOK, w.Code)

	//check that response from body is "successful"
	assert.Equal(t, "successful", w.Body.String())
}

func TestSignUp(t *testing.T){
	//creating the router
	router := gin.Default()

	//signup endpoint
	router.POST("/signup", func(c *gin.Context){

		var requestBody map[string]string
		if err := c.BindJSON(&requestBody); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		//checking if email and password are present
		if requestBody["email"] == "" || requestBody["password"] == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "email and password are required"})
			return
		}

		user := struct{
			Username string 'json:"username"'
			Password string 'json:"password"'
		}{
			Username: requestBody["username"],
			Password: requestBody["password"]
		}

	})

}
