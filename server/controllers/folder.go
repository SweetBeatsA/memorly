package controllers

import (
	"context"
	// "fmt"
	// "memorly/configs"
	"memorly/forms"
	//"memorly/helpers"
	"memorly/models"
	"memorly/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "golang.org/x/crypto/bcrypt"
)

func CreateFolder() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Your Implementation Should Go Here
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var folder forms.CreateFolderForm

		defer cancel()
		if err := c.BindJSON(&folder); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "Binding Error", Data: nil})
			return
		}

		if validationErr := validate.Struct(&folder); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "Validation Error", Data: nil})
			return
		}

		count, err := userCollection.CountDocuments(ctx, bson.M{"Title": folder.Title})
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "Error", Data: nil})
			return
		}

		if count > 0 {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "Folder Name Already Taken", Data: nil})
			return
		}

		newFolder := models.Folder{
			Id:        primitive.NewObjectID(),
			// CreatorId: primitive.ObjectID(),
			//^ the above isnt working not sure how get the jwt token the ID and Creator ID since they already exist
			//Also I might be wrong but I assume the ID is a new ID for the folder thats created so I made a new object ID for it
			Title:      folder.Title,
			//Not sure if I need to include the created time and updated time
		}

		newFolder.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		newFolder.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		//Again not sure if this is now the created and updated time

		//_, err = folderCollection.InsertOne(ctx, newFolder)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "Database Error", Data: nil})
			return
		}

		//Not sure what to check to make sure the folder was created
	}
}
