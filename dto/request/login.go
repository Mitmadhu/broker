package request

import (
	"net/http"

	"github.com/Mitmadhu/broker/helper"
)

type LoginRequest struct {
	BaseRequest
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	BaseRequest
	Username  string `json:"username"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       uint16 `json:"age"`
}

func (l LoginRequest) HasError(w http.ResponseWriter) bool {
	errs := []error{}
	helper.CheckEmpty(l.MsgId, &errs, "msg_id")
	helper.CheckEmpty(l.Password, &errs, "password")
	helper.CheckEmpty(l.Username, &errs, "username")
	if len(errs) > 0 {
		helper.SendErrorResponseArray(w, errs)
	}
	return len(errs) != 0
}

func (l RegisterRequest) HasError(w http.ResponseWriter) bool {
	errs := []error{}
	helper.CheckEmpty(l.MsgId, &errs, "msg_id")
	helper.CheckEmpty(l.Password, &errs, "password")
	helper.CheckEmpty(l.Username, &errs, "username")
	helper.CheckEmpty(l.FirstName, &errs, "first_name")
	helper.CheckEmpty(l.LastName, &errs, "last_name")

	if len(errs) > 0 {
		helper.SendErrorResponseArray(w, errs)
	}
	return len(errs) != 0
}
