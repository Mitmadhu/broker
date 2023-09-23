package helper

import (
	"encoding/json"
	"net/http"

	"github.com/Mitmadhu/broker/dto"
	"github.com/Mitmadhu/broker/dto/response"
)

func SendErrorResponse(w http.ResponseWriter, msg string, code uint64) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(int(code))
	errResp := response.ErrorResponse{
		MsgId:      "dummyvalue",
		Message:    msg,
		StatusCode: response.HttpStatus(code),
	}
	b, _ := json.Marshal(errResp)
	w.Write(b)
}

func SendSuccessResponse(w http.ResponseWriter, resp interface{}, code uint64) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(int(code))
	succResp := response.SuccessResponse{
		MsgId:      "dummy",
		StatusCode: response.HttpStatus(code),
		Response:   resp,
	}
	b, _ := json.Marshal(succResp)
	w.Write(b)
}

func SendSuccessRespWithClaims(w http.ResponseWriter, resp interface{}, code uint64, claims JWTValidation) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(int(code))
	succResp := response.SuccessResponse{
		MsgId:        "dummy",
		StatusCode:   response.HttpStatus(code),
		BaseResponse: dto.BaseResponse{
			IsTokenRefresh: claims.IsRefreshed,
			AccessToken: claims.AccessToken,
			RefreshToken: claims.RefreshToken,
		},
		Response:     resp,
	}
	b, _ := json.Marshal(succResp)
	w.Write(b)
}

func SendErrorResponseArray(w http.ResponseWriter, errs []error) {
	var messages string
	if len(errs) == 0 {
		messages = ""
	} else {
		messages = errs[0].Error()
	}

	for _, err := range errs[1:] {
		messages += "\n " + err.Error()
	}
	errResp := response.ErrorResponse{
		MsgId:      "dummy error response",
		StatusCode: response.HttpStatus(http.StatusBadRequest),
		Message:    messages,
	}
	b, _ := json.Marshal(errResp)
	w.Write(b)
}
