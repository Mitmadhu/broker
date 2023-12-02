package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	jwtAuth "github.com/Mitmadhu/broker/auth/jwt"
	"github.com/Mitmadhu/commons/constants"
	"github.com/Mitmadhu/commons/helper"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type ErrorHandler interface {
	HasError(w http.ResponseWriter) bool
}

type RouterRequest struct {
	DTO            ErrorHandler
	Method         string
	Handler        func(http.ResponseWriter, interface{})
	ValidationType string
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (t Token) HasError(w http.ResponseWriter) bool {
	return false
}

var RouterMap = map[string]RouterRequest{}

func Routers(port int) {
	r := mux.NewRouter()
	addApis(r)

	// Configure CORS middleware
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),                   // Replace with your allowed origins
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Adjust the allowed methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),           // Adjust the allowed headers
	)

	fmt.Printf("Server is listening on :%d...\n", port)
	http.Handle("/", corsOptions(r))
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		panic(err)
	}
}

func addApis(r *mux.Router) {
	for path, _ := range RouterMap {
		r.HandleFunc(path, middleHandler)
	}
}

func middleHandler(w http.ResponseWriter, r *http.Request) {
	b, err := io.ReadAll(r.Body)
	if err != nil {
		helper.SendErrorResponse(w, "invalid request body", "", http.StatusBadRequest)
		return
	}
	reqObj, ok := RouterMap[r.URL.Path]
	if !ok {
		helper.SendErrorResponse(w, "invalid URL", "", http.StatusNotFound)
	}
	json.Unmarshal(b, reqObj.DTO)
	// check for errors
	if reqObj.DTO.HasError(w) {
		return
	}
	if reqObj.ValidationType == constants.NoneValidation {
		reqObj.Handler(w, reqObj.DTO)
		return
	}

	// check for auth token
	req := &Token{}
	json.Unmarshal(b, req)
	switch reqObj.ValidationType {
	case constants.JWTValidation:
		// its not coming here
		if jwtAuth.IsJWTTokenExpired(req.AccessToken, req.RefreshToken) {
			helper.SendErrorResponse(w, "", constants.	TokenExipired, http.StatusUnauthorized)
			return
		}
	}
	reqObj.Handler(w, reqObj.DTO)

}
