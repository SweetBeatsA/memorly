package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine{
    router := gin.Default()

	router.GET("/getuser", func(c *gin.Context){
		c.String(http.StatusOK, "successful")
	})

    return router
};

func TestGetUser(t *testing.T){
	router := setupRouter()

	//request to /login
	request, err := http.NewRequest("GET", "/getuser", nil)
	assert.NoError(t, err)

	//record response
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	//checking if response 200 OK is ok
	assert.Equal(t, http.StatusOK, w.Code)

	//check that response from body is "successful"
	assert.Equal(t, "successful", w.Body.String())

}
