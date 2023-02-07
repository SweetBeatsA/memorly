package models

import (
	"errors"
	"gin/forms"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct{}

type User struct {
	Id       int    `json:"id" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = []User{
	{Id: 1, Name: "First User", Email: "first@gmail.com", Password: "firstPassword"},
	{Id: 2, Name: "Second User", Email: "second@gmail.com", Password: "secondPassword"},
	{Id: 3, Name: "Third User", Email: "third@gmail.com", Password: "thirdPassword"},
}

// var authModel = new(AuthModel)

// func login(form forms.LoginForm) (token authModel.Token, err error) {
// 	var token authModel.Token

// 	return token, nil
// }

func register(form forms.RegisterForm) (user User, err error) {
	bytePassword := []byte(form.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)

	if err != nil {
		return user, errors.New("Please check the password format")
	}

	user.Id = len(users) + 1
	user.Name = form.Name
	user.Email = form.Email
	user.Password = string(hashedPassword)

	return user, err
}

func getAll() (users []User, err error) {
	return users, err
}
