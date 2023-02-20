package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id         primitive.ObjectID `json:"id,omitempty"`
	Name       string             `json:"name,omitempty"`
	Email      string             `json:"email,omitempty"`
	Password   string             `json:"password,omitempty"`
	Created_at time.Time          `json:"created_at`
	Updated_at time.Time          `json:"updated_at`
}
