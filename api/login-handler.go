package api

import (
	"net/http"

	"github.com/Mitmadhu/broker/auth/jwt"
	"github.com/Mitmadhu/broker/constants"
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
		helper.SendErrorResponse(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if !validate{
		helper.SendErrorResponse(w, "invalid username/password", http.StatusUnauthorized)
		return
	}
	
	token, err := jwtAuth.GenerateToken(req.Username, constants.AccessToken)
	if err != nil{
		helper.SendErrorResponse(w, "internal server error", http.StatusInternalServerError)
		return
	}
	resp := response.LoginResponse{
		Success: true,
		Token: token,
		IsTokenRefreshed: true,
	}
	helper.SendSuccessResponse(w, resp, http.StatusAccepted)
}