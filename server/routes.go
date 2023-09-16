package server

import (
	"fmt"
	"net/http"

	"github.com/Mitmadhu/broker/api"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func Routers(){
	r := mux.NewRouter()

    // Define a route for handling GET requests to the root path "/"
    r.HandleFunc("/user", api.GetUserDetails).Methods("POST")

	 // Configure CORS middleware
	corsOptions := handlers.CORS(
        handlers.AllowedOrigins([]string{"http://localhost:3000"}), // Replace with your allowed origins
        handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}), // Adjust the allowed methods
        handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Adjust the allowed headers
    )

    port := 8080
    fmt.Printf("Server is listening on :%d...\n", port)
    http.Handle("/", corsOptions(r))
    err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
    if err != nil {
        panic(err)
    }
}



