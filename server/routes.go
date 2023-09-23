package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Mitmadhu/broker/api"
	"github.com/Mitmadhu/broker/dto/request"
	"github.com/Mitmadhu/broker/helper"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type ErrorHandler interface {
	HasError(w http.ResponseWriter) bool
}


type routerRequest struct{
	dto ErrorHandler
	method string
	handler func(http.ResponseWriter, interface{})
}

var routerMap = map[string]routerRequest{
	"/user-details": {
		dto: &request.UserDetailsRequest{},
		method: http.MethodGet,
		handler: api.GetUserDetails,
	},
	"/login": {
		dto: &request.LoginRequest{},
		method: http.MethodPost,
		handler: api.Login,
	},
	"/register" : {
		dto : &request.RegisterRequest{},
		method: http.MethodPost,
		handler: api.Register,
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
	for path, _:= range routerMap{
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
	if reqObj.dto.HasError(w){
		return
	}
	reqObj.handler(w, reqObj.dto)
}