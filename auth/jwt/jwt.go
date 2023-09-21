package jwtAuth

import (
	"errors"
	"time"

	"github.com/Mitmadhu/broker/constants"
	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("secret-madhu-ka-key")

type customClaims struct {
	jwt.StandardClaims
	Username    string `json:"username"`
	RefreshTime int64  `json:"refresh_time"`
	
}

func GenerateToken(username string, tokenType string) (string, error) {
	var expiryTime int64
	switch tokenType{
	case constants.AccessToken:
		expiryTime = time.Now().Add(time.Second).Unix()
	case constants.RefreshToken:
		expiryTime = time.Now().Add(time.Minute).Unix()
	default:
		panic("invalid tokentype to genereate jwt token")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims{	
		Username: username,
		RefreshTime: time.Now().Add(time.Hour).Unix(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiryTime,
		},
	})
	
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func Validate(tokenString string) (*customClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &customClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return nil, errors.New("errror converting standard claim")
	}
	return claims, nil
}
