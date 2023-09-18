package request


import (
	"net/http"
	"github.com/Mitmadhu/broker/helper"
)

type LoginRequest struct{
	Username string `json:"username"`
	Password string `json:"password"`
}

func (l LoginRequest) CheckError(w http.ResponseWriter) bool{
	errs := []error{}
	helper.CheckEmpty(l.Password, &errs, "password")
	helper.CheckEmpty(l.Username, &errs, "username")
	if len(errs) > 0 {
		helper.SendErrorResponseArray(w, errs)
	}
	return len(errs) != 0
}

func (l LoginRequest) Test(){
	
}