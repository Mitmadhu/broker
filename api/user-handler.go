package api

import (
	"net/http"

	"github.com/Mitmadhu/broker/dto/response"
	"github.com/Mitmadhu/broker/helper"
	"github.com/Mitmadhu/mysqlDB/config"
	"github.com/Mitmadhu/mysqlDB/database/model"
)

func GetUserDetails(w http.ResponseWriter, r *http.Request) {

	// validate token
	
	// get user details
	u := model.User{}
	user, err := u.GetUserByID(config.GetDB(), "1")
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
