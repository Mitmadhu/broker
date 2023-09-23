package helper

import (

	jwtAuth "github.com/Mitmadhu/broker/auth/jwt"
)

type JWTValidation struct {
	AccessToken  string
	RefreshToken string
	IsRefreshed  bool
	CustomClaims *jwtAuth.CustomClaims
}

func GetJWTClaims(aToken, rToken string) (*JWTValidation, error) {
	claim, err := jwtAuth.Validate(aToken)
	if err != nil {
		claim, err := jwtAuth.Validate(rToken)
		if err != nil {
			return nil, err
		}
		accessToken, refreshToken, err := jwtAuth.GenerateToken(claim.Username)
		if err != nil {
			panic(err.Error())
		}
		return &JWTValidation{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			IsRefreshed:  true,
			CustomClaims: claim,
		}, nil
	}
	return &JWTValidation{IsRefreshed: false, CustomClaims: claim}, nil
}
