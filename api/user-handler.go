package api

import (
	"net/http"

	"github.com/Mitmadhu/broker/dto/request"
	"github.com/Mitmadhu/broker/dto/response"
	"github.com/Mitmadhu/broker/helper"
	"github.com/Mitmadhu/mysqlDB/database/model"
)

func GetUserDetails(w http.ResponseWriter, dto interface{}) {
	req, ok := dto.(*request.UserDetailsRequest)
	if !ok {
		helper.SendErrorResponse(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// validate token
	claims, err := helper.GetJWTClaims(req.AccessToken, req.RefreshToken)
	if err != nil {
		helper.SendErrorResponse(w, "invalid token", http.StatusUnauthorized)
		return
	}
	 
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
	helper.SendSuccessRespWithClaims(w, resp, http.StatusAccepted, *claims)

}
