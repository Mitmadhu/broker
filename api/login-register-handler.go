package api

import (
	"fmt"
	"net/http"

	clients "github.com/Mitmadhu/commons/clients"

	jwtAuth "github.com/Mitmadhu/broker/auth/jwt"
	"github.com/Mitmadhu/broker/config"
	clients "github.com/Mitmadhu/commons/clients"

	jwtAuth "github.com/Mitmadhu/broker/auth/jwt"
	"github.com/Mitmadhu/broker/config"
	"github.com/Mitmadhu/broker/dto/request"
	"github.com/Mitmadhu/broker/dto/response"
	"github.com/Mitmadhu/commons/constants"
	cmnHelper "github.com/Mitmadhu/commons/helper"
	"github.com/Mitmadhu/commons/constants"
	cmnHelper "github.com/Mitmadhu/commons/helper"
	"github.com/Mitmadhu/mysqlDB/database/model"
	mysqlDto "github.com/Mitmadhu/mysqlDB/dto"
	mysqlDto "github.com/Mitmadhu/mysqlDB/dto"
)

func Login(w http.ResponseWriter, dto interface{}) {
	req, ok := dto.(*request.LoginRequest)
	if !ok {
		cmnHelper.SendErrorResponse(w, "", "invalid request body", http.StatusBadRequest)
		cmnHelper.SendErrorResponse(w, "", "invalid request body", http.StatusBadRequest)
		return
	}

	// TODO validate user
	mysqlClient := clients.MysqlClientImpl{}
	loginResp, err := mysqlClient.Login(config.Configs.Endpoints[constants.MYSQLDB], "POST", mysqlDto.ValidateUserRequest{
		Username: req.Username,
		Password: req.Password,
		BaseRequest: mysqlDto.BaseRequest{
			MsgID: req.MsgID,
		},
	})

	if err != nil {
		println(err.Error())
		cmnHelper.SendErrorResponse(w, req.MsgID, "", http.StatusInternalServerError)
		return
	}
	// TODO: validate loginResp
	if loginResp.IsValid == nil {
		println("nil login response received from mysql")
		cmnHelper.SendErrorResponse(w, req.MsgID, "", http.StatusInternalServerError)

	// TODO validate user
	mysqlClient := clients.MysqlClientImpl{}
	loginResp, err := mysqlClient.Login(config.Configs.Endpoints[constants.MYSQLDB], "POST", mysqlDto.ValidateUserRequest{
		Username: req.Username,
		Password: req.Password,
		BaseRequest: mysqlDto.BaseRequest{
			MsgID: req.MsgID,
		},
	})

	if err != nil {
		println(err.Error())
		cmnHelper.SendErrorResponse(w, req.MsgID, "", http.StatusInternalServerError)
		return
	}
	// TODO: validate loginResp
	if loginResp.IsValid == nil {
		println("nil login response received from mysql")
		cmnHelper.SendErrorResponse(w, req.MsgID, "", http.StatusInternalServerError)
		return
	}

	if !(*loginResp.IsValid) {
		cmnHelper.SendErrorResponse(w, req.MsgID, constants.UserNotFound, http.StatusNotFound)
	if !(*loginResp.IsValid) {
		cmnHelper.SendErrorResponse(w, req.MsgID, constants.UserNotFound, http.StatusNotFound)
		return
	}


	accessToken, refreshToken, err := jwtAuth.GenerateToken(req.Username)
	if err != nil {
	if err != nil {
		fmt.Printf("error while generating jwt token, err: %v", err.Error())
		cmnHelper.SendErrorResponse(w, req.MsgID, "", http.StatusInternalServerError)
		cmnHelper.SendErrorResponse(w, req.MsgID, "", http.StatusInternalServerError)
		return
	}
	resp := response.LoginResponse{
		BaseResponse: response.BaseResponse{
			MsgID:          req.MsgID,
			Success:        true,
			StatusCode:     http.StatusOK,
			AccessToken:    accessToken,
			RefreshToken:   refreshToken,
			IsTokenRefresh: true,
		},
		BaseResponse: response.BaseResponse{
			MsgID:          req.MsgID,
			Success:        true,
			StatusCode:     http.StatusOK,
			AccessToken:    accessToken,
			RefreshToken:   refreshToken,
			IsTokenRefresh: true,
		},
		Success: true,
	}

	cmnHelper.SendSuccessResponse(w, resp, http.StatusOK)

	cmnHelper.SendSuccessResponse(w, resp, http.StatusOK)
}

func Register(w http.ResponseWriter, dto interface{}) {
func Register(w http.ResponseWriter, dto interface{}) {
	req, ok := dto.(*request.RegisterRequest)
	if !ok {
		cmnHelper.SendErrorResponse(w, "", "invalid request body", http.StatusBadRequest)
		cmnHelper.SendErrorResponse(w, "", "invalid request body", http.StatusBadRequest)
		return
	}
	// check if username is available
	u := model.User{}
	_, err := u.GetUserByUsername(req.Username)
	if err == nil {
		cmnHelper.SendErrorResponse(w, req.MsgID, constants.UsernameExists, http.StatusBadRequest)
		cmnHelper.SendErrorResponse(w, req.MsgID, constants.UsernameExists, http.StatusBadRequest)
		return
	}

	err = model.User{}.Register(req.Username, req.Password, req.FirstName, req.LastName, req.Age)
	// model.User.Register(dto.u)
	if err != nil {
	if err != nil {
		println(err)
		cmnHelper.SendErrorResponse(w, req.MsgID, "something went wrong", http.StatusInternalServerError)
		return
	}
	// generate token
	accessToken, refreshToken := "", ""
	// generate token
	accessToken, refreshToken := "", ""
	response := response.RegisterResponse{
		BaseResponse: response.BaseResponse{
			MsgID:        req.MsgID,
			Success:      true,
			StatusCode:   http.StatusAccepted,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		BaseResponse: response.BaseResponse{
			MsgID:        req.MsgID,
			Success:      true,
			StatusCode:   http.StatusAccepted,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		},
		Username: req.Username,
	}
	cmnHelper.SendSuccessResponse(w, response, http.StatusCreated)
	cmnHelper.SendSuccessResponse(w, response, http.StatusCreated)
}

