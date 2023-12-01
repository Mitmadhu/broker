package server

import (
	"net/http"

	"github.com/Mitmadhu/broker/api"
	"github.com/Mitmadhu/broker/config"
	"github.com/Mitmadhu/broker/constants"
	"github.com/Mitmadhu/broker/dto/request"
	router "github.com/Mitmadhu/commons/server"
)
func InitRouter(){
	router.RouterMap = map[string]router.RouterRequest{
		"/user-details": {
			DTO:          &request.UserDetailsRequest{},
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
		"/register" : {
			DTO : &request.RegisterRequest{},
			Method: http.MethodPost,
			Handler: api.Register,
			ValidationType: constants.NoneValidation,
		},
	}
	router.Routers(config.Configs.Port)
}