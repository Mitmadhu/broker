package api

import (
	"fmt"
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
		helper.SendErrorResponse(w, "", "invalid request body", http.StatusBadRequest)
		return
	}
	// validate password
	u := model.User{}
	validate, err := u.ValidateUser(req.Username, req.Password)
	
	if err != nil{
		helper.SendErrorResponse(w, req.MsgId, err.Error(), http.StatusUnauthorized)
		return
	}

	if !validate{
		helper.SendErrorResponse(w, req.MsgId, "invalid username/password", http.StatusUnauthorized)
		return
	}
	
	accessToken, refreshToken, err := jwtAuth.GenerateToken(req.Username)
	if err != nil{
		fmt.Printf("error while generating jwt token, err: %v", err.Error())
		helper.SendErrorResponse(w, req.MsgId, "internal server error", http.StatusInternalServerError)
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
	helper.SendSuccessRespWithClaims(w, req.MsgId, resp, http.StatusAccepted, claims)
}

func Register(w http.ResponseWriter, dto interface{}){
	req, ok := dto.(*request.RegisterRequest)
	if !ok {
		helper.SendErrorResponse(w, "", "invalid request body", http.StatusBadRequest)
		return
	}
	// check if username is available
	u := model.User{}
	_, err := u.GetUserByUsername(req.Username)
	if err == nil {
		helper.SendErrorResponse(w, req.MsgId, constants.UsernameExists, http.StatusBadRequest)
		return
	}

	err = model.User{}.Register(req.Username, req.Password, req.FirstName, req.LastName, req.Age)
	// model.User.Register(dto.u)
	if (err != nil){
		println(err)
		helper.SendErrorResponse(w, req.MsgId, "something went wrong", http.StatusInternalServerError )
		return
	}
	response := response.RegisterResponse{
		Username: req.Username,
	}
	helper.SendSuccessResponse(w, req.MsgId,response, http.StatusCreated)
}