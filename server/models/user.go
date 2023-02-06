package models

type User struct {
	id   int    `json:"id" gorm:"primary_key"`
	name string `json:"name"`
	email string `json:"email"`
	password string `json:"password"`
}