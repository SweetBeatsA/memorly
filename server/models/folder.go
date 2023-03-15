package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Folder struct {
	Id        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	CreatorId primitive.ObjectID `json:"creatorId" bson:"creatorId,omitempty"`
	Title     string             `json:"title" bson:"title,omitempty"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
}
