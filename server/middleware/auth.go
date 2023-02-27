package middleware

import (
	"fmt"
	"memorly/helpers"
	"memorly/responses"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: fmt.Sprintf("No Authorization Token Provided"), Data: nil})
			c.Abort()
			return
		}

		signature, err := helpers.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusUnauthorized, responses.Response{Status: http.StatusUnauthorized, Message: err, Data: nil})
			c.Abort()
			return
		}
		c.Set("email", signature.Email)
		c.Set("name", signature.Name)
		c.Set("id", signature.Id)
		c.Next()
	}
}
