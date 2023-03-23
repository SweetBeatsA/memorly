package controllers

import (
	"context"
	"memorly/configs"
	"memorly/forms"
	"memorly/models"
	"memorly/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var cardCollection *mongo.Collection = configs.GetCollection(configs.DB, "cards")

func CreateCard() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var card forms.CreateCardForm

		defer cancel()
		if err := c.BindJSON(&card); err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "Binding Error", Data: nil})
			return
		}

		if validationErr := validate.Struct(&card); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusBadRequest, Message: "Validation Error", Data: nil})
			return
		}

		var folder models.Folder

		err := folderCollection.FindOne(ctx, bson.M{"_id": card.FolderId}).Decode(&folder)
		defer cancel()

		if err != nil {
			c.JSON(http.StatusBadRequest, responses.Response{Status: http.StatusNotFound, Message: "No Matched Folder", Data: nil})
			return
		}

		id, _ := c.Get("id")
		strId, _ := id.(string)
		creatorId, _ := primitive.ObjectIDFromHex(strId)

		newCard := models.Card{
			Id:        primitive.NewObjectID(),
			FolderId:  folder.Id,
			Question:  card.Question,
			Answer:    card.Answer,
			CreatorId: creatorId,
		}

		newCard.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		newCard.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		_, err = cardCollection.InsertOne(ctx, newCard)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.Response{Status: http.StatusInternalServerError, Message: "Database Error", Data: nil})
			return
		}

		c.JSON(http.StatusCreated, responses.Response{Status: http.StatusCreated, Message: "Success", Data: map[string]interface{}{"id": newCard.Id}})
	}
}
