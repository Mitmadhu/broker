package api

import (
	"fmt"
	"net/http"

	clients "github.com/Mitmadhu/commons/clients"
	jwtAuth "github.com/Mitmadhu/broker/auth/jwt"
	"github.com/Mitmadhu/broker/config"
	"github.com/Mitmadhu/broker/dto/request"
	"github.com/Mitmadhu/broker/dto/response"
	"github.com/Mitmadhu/commons/constants"
	cmnHelper "github.com/Mitmadhu/commons/helper"
	mysqlDto "github.com/Mitmadhu/mysqlDB/dto"
)

func Login(w http.ResponseWriter, dto interface{}) {
	req, ok := dto.(*request.LoginRequest)
	if !ok {
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

	// TODO: validate loginResp
	if loginResp.IsValid == nil {
		println("nil login response received from mysql")
		cmnHelper.SendErrorResponse(w, req.MsgID, "", http.StatusInternalServerError)
		return
	}
	if err != nil {
		println(err.Error())
		cmnHelper.SendErrorResponse(w, req.MsgID, "", http.StatusInternalServerError)
		return
	}

	if !(*loginResp.IsValid) {
		cmnHelper.SendErrorResponse(w, req.MsgID, constants.UserNotFound, http.StatusNotFound)
		return
	}

	accessToken, refreshToken, err := jwtAuth.GenerateToken(req.Username)
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
		Success: true,
	}

	cmnHelper.SendSuccessResponse(w, resp, http.StatusOK)

}

func Register(w http.ResponseWriter, dto interface{}) {
	req, ok := dto.(*request.RegisterRequest)
	if !ok {
		cmnHelper.SendErrorResponse(w, "", "invalid request body", http.StatusBadRequest)
		return
	}
	
	// register http call
	mysqlClient := clients.MysqlClientImpl{}
	registerResp, err := mysqlClient.Register(config.Configs.Endpoints[constants.MYSQLDB], http.MethodGet, mysqlDto.RegisterUserRequest{
		BaseRequest: mysqlDto.BaseRequest{
			MsgID: req.MsgID,
		},
		Username: req.Username,
		Password: req.Password,
		LastName: req.LastName,
		FirstName: req.FirstName,
		Age : req.Age,
	})

	if (err != nil){
		println(err.Error())
		cmnHelper.SendErrorResponse(w, req.MsgID, "", http.StatusInternalServerError)
		return
	}
	
	if registerResp.Message == constants.UsernameExists {
		cmnHelper.SendErrorResponse(w, req.MsgID, constants.UsernameExists, uint64(registerResp.StatusCode))
		return
	}

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
		Username: req.Username,
	}
	cmnHelper.SendSuccessResponse(w, response, http.StatusCreated)
}
