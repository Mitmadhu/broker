package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mitmadhu/commons/handle_http"

	"github.com/Mitmadhu/broker/config"
	"github.com/Mitmadhu/broker/auth/jwt"
	"github.com/Mitmadhu/broker/constants"
	"github.com/Mitmadhu/broker/dto/request"
	"github.com/Mitmadhu/broker/dto/response"
	cmnConsts "github.com/Mitmadhu/commons/constants"
	cmnHelper "github.com/Mitmadhu/commons/helper"
	"github.com/Mitmadhu/broker/helper"
	"github.com/Mitmadhu/mysqlDB/database/model"
	mysqlDto "github.com/Mitmadhu/mysqlDB/dto"
)

func Login(w http.ResponseWriter, dto interface{}) {
	req, ok := dto.(*request.LoginRequest)
	if !ok {
		cmnHelper.SendErrorResponse(w, "", "invalid request body", http.StatusBadRequest)
		return
	}
	// validate password
	byteResp, err := handle_http.Call(config.Configs.Endpoints[cmnConsts.MYSQLDB] + "/user-exists", http.MethodPost, mysqlDto.ValidateUserRequest{
		BaseRequest: mysqlDto.BaseRequest{MsgId: req.MsgId},
		Username:    req.Username,
		Password:    req.Password,
	})

	if err != nil {
		cmnHelper.SendErrorResponse(w, req.MsgId, err.Error(), http.StatusUnauthorized)
		return
	}
	validate := &mysqlDto.ValidateUserResponse{}
	json.Unmarshal(byteResp, validate)
	
	if !validate.IsValid {
		cmnHelper.SendErrorResponse(w, req.MsgId, constants.UserNotFound , http.StatusUnauthorized)
		return
	}

	accessToken, refreshToken, err := jwtAuth.GenerateToken(req.Username)
	if err != nil{
		fmt.Printf("error while generating jwt token, err: %v", err.Error())
		cmnHelper.SendErrorResponse(w, req.MsgId, "internal server error", http.StatusInternalServerError)
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

func Register(w http.ResponseWriter, dto interface{}) {
	req, ok := dto.(*request.RegisterRequest)
	if !ok {
		cmnHelper.SendErrorResponse(w, "", "invalid request body", http.StatusBadRequest)
		return
	}
	// check if username is available
	u := model.User{}
	_, err := u.GetUserByUsername(req.Username)
	if err == nil {
		cmnHelper.SendErrorResponse(w, req.MsgId, constants.UsernameExists, http.StatusBadRequest)
		return
	}

	err = model.User{}.Register(req.Username, req.Password, req.FirstName, req.LastName, req.Age)
	// model.User.Register(dto.u)
	if err != nil {
		println(err)
		cmnHelper.SendErrorResponse(w, req.MsgId, "something went wrong", http.StatusInternalServerError)
		return
	}
	response := response.RegisterResponse{
		Username: req.Username,
	}
	cmnHelper.SendSuccessResponse(w, req.MsgId, response, http.StatusCreated)
}
