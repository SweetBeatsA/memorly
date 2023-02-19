package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Name     string             `json:"name,omitempty"`
	Email    string             `json:"email,omitempty"`
	Password []byte             `json:"password,omitempty"`
}
