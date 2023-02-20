package helpers

import (
	"fmt"
	"log"
	"memorly/configs"
	"memorly/models"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Signature struct {
	Email string
	Name  string
	Id    primitive.ObjectID
	jwt.StandardClaims
}

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var SECRET_KEY string = os.Getenv("SECRET_KEY")

func GenerateAllTokens(user models.User) (accessToken string, refreshToken string, err error) {
	signature := Signature{
		Email: user.Email,
		Name:  user.Name,
		Id:    user.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(240)).Unix(),
		},
	}

	refreshSignature := Signature{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(1680)).Unix(),
		},
	}

	acToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, signature).SignedString([]byte(SECRET_KEY))
	reToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshSignature).SignedString([]byte(SECRET_KEY))

	if err != nil {
		log.Panic(err)
		return
	}

	return acToken, reToken, err
}

func ValidateToken(accessToken string) (signature *Signature, msg string) {
	token, err := jwt.ParseWithClaims(
		accessToken,
		&Signature{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	signature, ok := token.Claims.(*Signature)
	if !ok {
		msg = fmt.Sprintf("the token is invalid")
		msg = err.Error()
		return
	}

	if signature.ExpiresAt < time.Now().Local().Unix() {
		msg = fmt.Sprintf("token is expired")
		msg = err.Error()
		return
	}
	return signature, msg
}
