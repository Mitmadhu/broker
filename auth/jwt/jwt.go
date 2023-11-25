package jwtAuth

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secret-madhu-ka-key")

type CustomClaims struct {
	jwt.StandardClaims
	Username    string `json:"username"`
	
}

func generateJWTToken(username string, expiryTime int64) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{	
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime,
		},
	})
	
	return token.SignedString(secretKey)
}

func GenerateToken(username string) (string, string, error) {
	accessToken, err := generateJWTToken(username, time.Now().Add(time.Hour).Unix())
	if err != nil {
		return "", "", err
	}
	refreshToken, err := generateJWTToken(username, time.Now().Add(time.Hour * 6).Unix())
	if err != nil {
		return "", "", err
	}
	return accessToken, refreshToken, nil
}

func IsJWTTokenExpired(aToken, rToken string) bool{
	_, err := Validate(aToken)
	if err == nil {
		return false
	}
	_, err = Validate(rToken)
	return err != nil
}

func Validate(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("errror converting standard claim")
	}
	return claims, nil
}
