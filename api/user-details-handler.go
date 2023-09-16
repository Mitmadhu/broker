package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Mitmadhu/broker/dto/request"
	"github.com/Mitmadhu/broker/dto/response"
)

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
    // Get the current working directory
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid input"))
		return
	}
	req := request.UserRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("invalid input"))
		return
	}

	if req.Username != "ayush.madhu" || req.Password != "123"{
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("username/password is incorrect"))
		return
		
	}
	w.Header().Set("Content-Type", "application/json")
	resp := response.UserResponse{
		Username: "mad.madhu",
		Age: 32,
		Address: "samastipura",
		Email: "ayz.gmail@com",
	}
	b, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error while marshalling response"))
		return
	}
	w.Write([]byte(b))
}