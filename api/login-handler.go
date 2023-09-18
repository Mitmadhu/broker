package api

import (
	"net/http"
	"github.com/Mitmadhu/broker/dto/request"
	"github.com/Mitmadhu/broker/dto/response"
	"github.com/Mitmadhu/broker/helper"
	"github.com/Mitmadhu/broker/constants"
)

func Login(w http.ResponseWriter, r *http.Request) {
	obj := r.Context().Value(constants.ReqPtr)
	req, ok := obj.(*request.LoginRequest)
	if !ok {
		helper.SendErrorResponse(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// validate password
	if req.Username != "ayush.madhu" || req.Password != "123"{
		helper.SendErrorResponse(w, "invalid username/password", http.StatusUnauthorized)
		return
	}

	resp := response.LoginResponse{
		Success: true,
		Token: "dummyToken",
	}
	helper.SendSuccessResponse(w, resp, http.StatusAccepted)
}