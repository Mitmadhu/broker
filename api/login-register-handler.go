package api

import (
	"fmt"
	"net/http"

	clients "command-line-arguments/home/madhu/go/src/github.com/Mitmadhu/commons/clients/mysql_client_impl.go"

	jwtAuth "github.com/Mitmadhu/broker/auth/jwt"
	"github.com/Mitmadhu/broker/constants"
	"github.com/Mitmadhu/broker/dto/request"
	"github.com/Mitmadhu/broker/dto/response"
	"github.com/Mitmadhu/broker/helper"
	cmnHelper "github.com/Mitmadhu/commons/helper"
	"github.com/Mitmadhu/mysqlDB/database/model"
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
	resp, err := mysqlClient.Login("http://localhost:8081", "POST", mysqlDto.ValidateUserRequest{
		Username: "ayush",
		Password: "123",
		BaseRequest: mysqlDto.BaseRequest{
			MsgID: "123",
		},
	})

	if err != nil {
		return nil, err
	}


	accessToken, refreshToken, err := jwtAuth.GenerateToken(req.Username)
	if err != nil {
		fmt.Printf("error while generating jwt token, err: %v", err.Error())
		cmnHelper.SendErrorResponse(w, req.MsgId, "internal server error", http.StatusInternalServerError)
		return
	}
	resp := response.LoginResponse{
		Success: true,
	}
	claims := helper.JWTValidation{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		IsRefreshed:  true,
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
