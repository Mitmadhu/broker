package dto

import (
	"net/http"

	"github.com/Mitmadhu/commons/helper"
)

type ValidateUserRequest struct {
	BaseRequest
	Username string `json:"username"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

type ValidateUserResponse struct {
	IsValid bool `json:"is_valid"`
}

func (v ValidateUserRequest) HasError(w http.ResponseWriter) bool {
	errs := []error{}
	helper.CheckEmpty(v.MsgId, &errs, "msg_id")
	helper.CheckEmpty(v.Username, &errs, "username")
	if len(errs) > 0 {
		helper.SendErrorResponseArray(w, errs)
	}
	return len(errs) != 0
}
