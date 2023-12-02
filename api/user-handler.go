package api

import (
	"net/http"

	"github.com/Mitmadhu/broker/dto/request"
	"github.com/Mitmadhu/broker/dto/response"
	"github.com/Mitmadhu/broker/helper"
	cmnHelper "github.com/Mitmadhu/commons/helper"
	commonHelper "github.com/Mitmadhu/commons/helper"
	"github.com/Mitmadhu/mysqlDB/database/model"
)

func GetUserDetails(w http.ResponseWriter, dto interface{}) {
	req, ok := dto.(*request.UserDetailsRequest)
	if !ok {
		commonHelper.SendErrorResponse(w, "", "invalid request body", http.StatusBadRequest)
		return
	}
	// validate token
	claims, _ := helper.GetJWTClaims(req.AccessToken, req.RefreshToken)

	// get user details
	u := model.User{}
	user, err := u.GetUserByID("1")
	if err != nil {
		commonHelper.SendErrorResponse(w, "internal server err", req.MsgID, http.StatusInternalServerError)
		return
	}
	resp := response.UserDetailsResponse{
		BaseResponse: response.BaseResponse{
			MsgID:          req.MsgID,
			StatusCode:     http.StatusAccepted,
			Success:        true,
			IsTokenRefresh: claims.IsRefreshed,
			AccessToken:    claims.AccessToken,
			RefreshToken:   claims.RefreshToken,
		},
		Username: user.Username,
		Age:      user.Age,
		Address:  "bandal",
	}
	cmnHelper.SendSuccessResponse(w, resp, http.StatusAccepted)
}
