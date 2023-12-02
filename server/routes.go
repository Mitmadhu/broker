package server

import (
	"net/http"

	"github.com/Mitmadhu/broker/api"
	"github.com/Mitmadhu/broker/config"
	"github.com/Mitmadhu/broker/dto/request"
	"github.com/Mitmadhu/commons/constants"
	"github.com/Mitmadhu/commons/server"
)

func InitRouter() {
	server.RouterMap = map[string]server.RouterRequest{
		"/user-details": {
			DTO:            &request.UserDetailsRequest{},
			Method:         http.MethodGet,
			Handler:        api.GetUserDetails,
			ValidationType: constants.JWTValidation,
		},
		"/login": {
			DTO:            &request.LoginRequest{},
			Method:         http.MethodPost,
			Handler:        api.Login,
			ValidationType: constants.NoneValidation,
		},
		"/register": {
			DTO:            &request.RegisterRequest{},
			Method:         http.MethodPost,
			Handler:        api.Register,
			ValidationType: constants.NoneValidation,
		},
	}
	server.Routers(config.Configs.Port)
}
