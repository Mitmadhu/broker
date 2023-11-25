package helper

import (
	"encoding/json"
	"net/http"

	"github.com/Mitmadhu/broker/dto/response"
)

func SendErrorResponse(w http.ResponseWriter, msgID, msg string, code uint64) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(int(code))
	errResp := response.ErrorResponse{
		MsgId:      msgID,
		Message:    msg,
		StatusCode: response.HttpStatus(code),
		Success: false,
	}
	b, _ := json.Marshal(errResp)
	w.Write(b)
}

func SendSuccessResponse(w http.ResponseWriter, msgID string, resp interface{}, code response.HttpStatus) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(int(code))
	succResp := response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			MsgID: msgID,
			StatusCode: code,
			Success: true,
		},
		Response:   resp,
	}
	b, _ := json.Marshal(succResp)
	w.Write(b)
}

func SendSuccessRespWithClaims(w http.ResponseWriter, msgID string, resp interface{}, code response.HttpStatus, claims JWTValidation) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(int(code))
	succResp := response.SuccessResponse{
		BaseResponse: response.BaseResponse{
			MsgID: msgID,
			StatusCode: code,
			Success: true,
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
		MsgId:      "",
		StatusCode: response.HttpStatus(http.StatusBadRequest),
		Message:    messages,
	}
	b, _ := json.Marshal(errResp)
	w.Write(b)
}
