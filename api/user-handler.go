package api

import (
	"net/http"

	"github.com/Mitmadhu/broker/dto/response"
	"github.com/Mitmadhu/broker/helper"
	"github.com/Mitmadhu/mysqlDB/database/model"
)

func GetUserDetails(w http.ResponseWriter, dto interface{}) {

	// validate token
	println("user-details")
	// get user details
	u := model.User{}
	user, err := u.GetUserByID("1")
	if err != nil{
		helper.SendErrorResponse(w, "internal server err", http.StatusInternalServerError)
	}
	resp := response.UserDetailsResponse{
		Username: user.Username,
		Age:      user.Age,
		Address:  "bandal",
	}
	helper.SendSuccessResponse(w, resp, http.StatusAccepted)

}
