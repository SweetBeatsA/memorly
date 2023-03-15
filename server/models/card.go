package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatorId primitive.ObjectID `json:"creatorId" bson:"creatorId,omitempty"`
	FolderId  primitive.ObjectID `json:"folderId" bson:"folderId,omitempty"`
	Question  string             `json:"question" bson:"question,omitempty"`
	Answer    string             `json:"answer" bson:"answer,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
}
