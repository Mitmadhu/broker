package helper

import (
	"encoding/json"
	"net/http"

	"github.com/Mitmadhu/broker/dto/response"
)

func SendErrorResponse(w http.ResponseWriter, msg string, code uint64){
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(int(code))
	errResp := response.ErrorResponse{
		MsgId: "dummyvalue",
		Message: msg,
		StatusCode: response.HttpStatus(code),
	}
	b, _ := json.Marshal(errResp)
	w.Write(b)
}

func SendSuccessResponse(w http.ResponseWriter, resp interface{}, code uint64){
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(int(code))
	succResp := response.SuccessResponse{
		MsgId: "dummy",
		StatusCode: response.HttpStatus(code),
		Response: resp,
	}
	b, _ := json.Marshal(succResp)
	w.Write(b)
}