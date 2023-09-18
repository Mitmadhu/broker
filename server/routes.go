package server

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Mitmadhu/broker/api"
	"github.com/Mitmadhu/broker/dto/request"
	"github.com/Mitmadhu/broker/helper"
	"github.com/Mitmadhu/broker/utils"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type ErrorHandler interface{
	CheckError(w http.ResponseWriter) bool
}

func Routers() {
	r := mux.NewRouter()

	// Define a route for handling GET requests to the root path "/"
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

func addApis(r *mux.Router){
	rapper(r, "/login", api.Login, "POST", &request.LoginRequest{})

	// user request
	rapper(r, "/user-details", api.GetUserDetails, "GET", &request.UserDetailsRequest{})
}

func rapper(r *mux.Router, path string, handler func(http.ResponseWriter, *http.Request), method string, reqObj ErrorHandler) {
    customMiddleware := setRequestObjType(reqObj)
	r.Use(customMiddleware)
	r.HandleFunc(path, handler)
    
}


// Middleware that sets a value in the request context

// Custom middleware that accepts arguments
func setRequestObjType(reqObj ErrorHandler) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            b, err := io.ReadAll(r.Body)
            if err != nil {
                helper.SendErrorResponse(w, "invalid request body", http.StatusBadRequest)
                return
            }
            json.Unmarshal(b, reqObj)
			if reqObj.CheckError(w){
				return
			}
            ctx := context.WithValue(r.Context(), utils.ReqPtr, reqObj)
            r = r.WithContext(ctx)
            // Call the next handler
            next.ServeHTTP(w, r)
        })
    }
}
