package auth

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"reddit/pkg/errors"
	"reddit/pkg/user"
	"time"
)

var (
	SigningToken      interface{} = SigningTokenValue
	SigningTokenValue             = []byte("2816e66cb08c9f4cb5d7c080b2fca85f17cdb1cbe32380c7fdde9cf469185e30")
	TokenKey                      = "Just token key"
	GenerateTokenFunc             = Generate
)

type AccessToken struct {
	Token string `json:"token"`
}

func Generate(userData user.User) (*AccessToken, *errors.Error) {
	userData.Password = ""
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": userData,
		"iat":  time.Now().Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})
	tokenString, tokenError := token.SignedString(SigningToken)
	if tokenError != nil {
		log.Printf("Auth: Generate error: %s\n", tokenError.Error())
		return nil, errors.New(http.StatusInternalServerError, errors.InternalError)
	}
	return &AccessToken{Token: tokenString}, nil
}
