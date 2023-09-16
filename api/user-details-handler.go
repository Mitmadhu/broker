package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Mitmadhu/broker/dto/request"
	"github.com/Mitmadhu/broker/dto/response"
	"github.com/Mitmadhu/broker/helper"
)

func GetUserDetails(w http.ResponseWriter, r *http.Request) {
    // Get the current working directory
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		helper.SendErrorResponse(w, "bad request body", http.StatusBadRequest)
		return
	}
	req := request.UserRequest{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		helper.SendErrorResponse(w, "bad request body", http.StatusBadRequest)
		return
	}

	if req.Username != "ayush.madhu" || req.Password != "123"{
		helper.SendErrorResponse(w, "username/password is incorrect", http.StatusUnauthorized)
		return
		
	}
	resp := response.UserResponse{
		Username: "mad.madhu",
		Age: 32,
		Address: "samastipura",
		Email: "ayz.gmail@com",
	}
	helper.SendSuccessResponse(w, resp, http.StatusAccepted)
}