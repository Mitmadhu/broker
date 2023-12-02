package helper

import (
	"encoding/json"
	"net/http"

	"github.com/Mitmadhu/commons/constants"
	"github.com/Mitmadhu/commons/dto/response"
)

func SendErrorResponse(w http.ResponseWriter, msgID, msg string, code uint64) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(int(code))
	
	if code == http.StatusInternalServerError {
		msg = constants.StatusInternalServerError
	}
	errResp := response.ErrorResponse{
		MsgId:      msgID,
		Message:    msg,
		StatusCode: response.HttpStatus(code),
		Success: false,
	}
	b, _ := json.Marshal(errResp)
	w.Write(b)
}

func SendSuccessResponse(w http.ResponseWriter, resp interface{}, code int) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	b, _ := json.Marshal(resp)
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
