package api

import (
	"fmt"
	"net/http"

	"github.com/Mitmadhu/broker/auth/jwt"
	"github.com/Mitmadhu/broker/dto/request"
	"github.com/Mitmadhu/broker/dto/response"
	"github.com/Mitmadhu/broker/helper"
	"github.com/Mitmadhu/mysqlDB/database/model"
)

func Login(w http.ResponseWriter, dto interface{}) {
	req, ok := dto.(*request.LoginRequest)
	if !ok {
		helper.SendErrorResponse(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// validate password
	u := model.User{}
	validate, err := u.ValidateUser(req.Username, req.Password)
	
	if err != nil{
		helper.SendErrorResponse(w, err.Error(), http.StatusUnauthorized)
		return
	}

	if !validate{
		helper.SendErrorResponse(w, "invalid username/password", http.StatusUnauthorized)
		return
	}
	
	accessToken, refreshToken, err := jwtAuth.GenerateToken(req.Username)
	if err != nil{
		fmt.Printf("error while generating jwt token, err: %v", err.Error())
		helper.SendErrorResponse(w, "internal server error", http.StatusInternalServerError)
		return
	}
	resp := response.LoginResponse{
		Success: true,
	}
	claims := helper.JWTValidation{
		AccessToken: accessToken,
		RefreshToken: refreshToken,
		IsRefreshed: true,
	}
	helper.SendSuccessRespWithClaims(w, resp, http.StatusAccepted, claims)
}