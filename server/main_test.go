package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	router := setupRouter()

	//request to /user
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
