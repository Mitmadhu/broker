package request

import (
	"net/http"

	"github.com/Mitmadhu/broker/helper"
)

type UserDetailsRequest struct {
	Username     string `json:"username"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (u *UserDetailsRequest) CheckError(w http.ResponseWriter) bool {
	errs := []error{}

	helper.CheckEmpty(u.Username, &errs, "username")
	helper.CheckEmpty(u.AccessToken, &errs, "access token")
	helper.CheckEmpty(u.RefreshToken, &errs, "refresh token")

	if len(errs) > 0 {
		helper.SendErrorResponseArray(w, errs)
	}
	return len(errs) != 0
}

func (u *UserDetailsRequest) Validate() bool {
	return false
}
