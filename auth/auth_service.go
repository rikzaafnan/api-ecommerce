package auth

import (
	"api-ecommerce/config"
	"api-ecommerce/helper"
	"errors"
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

type ServiceAuth interface {
	GenerateToken(userID int, role string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}

type jwtService struct {
}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userID int, role string) (string, error) {

	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	claim["random-string"] = helper.RandomString(25)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	signedToken, err := token.SignedString([]byte(config.LoadENV().JWTTOKEN))
	if err != nil {
		fmt.Println(err)
		return signedToken, err

	}

	return signedToken, nil

}

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {

	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}

		// return []byte(SECRET_KEY), nil
		return []byte(config.LoadENV().JWTTOKEN), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil

}
