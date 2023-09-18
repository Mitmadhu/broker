package request

import (
	"net/http"

	"github.com/Mitmadhu/broker/helper"
)

type UserDetailsRequest struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (u *UserDetailsRequest) CheckError(w http.ResponseWriter) bool {
	errs := []error{}

	helper.CheckEmpty(u.Username, &errs, "username")
	helper.CheckEmpty(u.Token, &errs, "token")

	if len(errs) > 0 {
		helper.SendErrorResponseArray(w, errs)
	}
	return len(errs) != 0
}
