package controllers

import (
	"context"
	"fmt"
	"log"
	"memorly/configs"
	"memorly/forms"
	"memorly/helpers"
	"memorly/models"
	"memorly/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 11)
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}

func VerifyPassword(userPassword string, providedPassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(providedPassword), []byte(userPassword))
	check := true
	msg := ""

	if err != nil {
		msg = fmt.Sprintf("Password Is Incorrect")
		check = false
	}
	return check, msg
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user forms.RegisterForm

		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if count > 0 {
			msg := "Email Already Taken"
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: msg, Data: map[string]interface{}{"data": msg}})
			return
		}

		password := HashPassword(user.Password)

		newUser := models.User{
			Id:       primitive.NewObjectID(),
			Name:     user.Name,
			Email:    user.Email,
			Password: password,
		}

		newUser.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		newUser.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		_, err = userCollection.InsertOne(ctx, newUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		accessToken, refreshToken, _ := helpers.GenerateAllTokens(newUser)
		c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"accessToken": accessToken, "refreshToken": refreshToken}})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user forms.LoginForm
		var foundUser models.User

		defer cancel()
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		err := userCollection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&foundUser)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "No User Matched With Email", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		valid, msg := VerifyPassword(user.Password, foundUser.Password)

		if valid == false {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: msg, Data: map[string]interface{}{"data": msg}})
			return
		}

		accessToken, refreshToken, _ := helpers.GenerateAllTokens(foundUser)

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"accessToken": accessToken, "refreshToken": refreshToken}})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var user models.User

		id, _ := c.Get("id")
		err := userCollection.FindOne(ctx, bson.M{"id": id}).Decode(&user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.Response{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}})
	}
}
