package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Mitmadhu/broker/api"
	jwtAuth "github.com/Mitmadhu/broker/auth/jwt"
	"github.com/Mitmadhu/broker/constants"
	"github.com/Mitmadhu/broker/dto/request"
	"github.com/Mitmadhu/broker/helper"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type ErrorHandler interface {
	HasError(w http.ResponseWriter) bool
}

type routerRequest struct {
	dto            ErrorHandler
	method         string
	handler        func(http.ResponseWriter, interface{})
	validationType string
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (t Token) HasError(w http.ResponseWriter) bool {
	return false
}

var routerMap = map[string]routerRequest{
	"/user-details": {
		dto:            &request.UserDetailsRequest{},
		method:         http.MethodGet,
		handler:        api.GetUserDetails,
		validationType: constants.JWTValidation,
	},
	"/login": {
		dto:            &request.LoginRequest{},
		method:         http.MethodPost,
		handler:        api.Login,
		validationType: constants.NoneValidation,
	},
	"/register": {
		dto:            &request.RegisterRequest{},
		method:         http.MethodPost,
		handler:        api.Register,
		validationType: constants.NoneValidation,
	},
}

func Routers() {
	r := mux.NewRouter()
	addApis(r)

	// Configure CORS middleware
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),                   // Replace with your allowed origins
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Adjust the allowed methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),           // Adjust the allowed headers
	)

	port := 8080
	fmt.Printf("Server is listening on :%d...\n", port)
	http.Handle("/", corsOptions(r))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}

func addApis(r *mux.Router) {
	for path, _ := range routerMap {
		r.HandleFunc(path, middleHandler)
	}
}

func middleHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		helper.SendErrorResponse(w, "invalid request body", "", http.StatusBadRequest)
		return
	}
	reqObj, ok := routerMap[r.URL.Path]
	if !ok {
		helper.SendErrorResponse(w, "invalid URL", "", http.StatusNotFound)
	}
	json.Unmarshal(b, reqObj.dto)
	// check for errors
	if reqObj.dto.HasError(w) {
		return
	}
	if reqObj.validationType == constants.NoneValidation {
		reqObj.handler(w, reqObj.dto)
		return
	}

	// check for auth token
	req := &Token{}
	json.Unmarshal(b, req)
	switch reqObj.validationType {
	case constants.JWTValidation:
		// its not coming here
		if jwtAuth.IsJWTTokenExpired(req.AccessToken, req.RefreshToken) {
			helper.SendErrorResponse(w, "", constants.TokenExipired, http.StatusUnauthorized)
			return
		}
	}
	reqObj.handler(w, reqObj.dto)

}