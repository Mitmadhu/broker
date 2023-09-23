package request

import (
	"net/http"

	"github.com/Mitmadhu/broker/helper"
)

type BaseRequest struct {
	MsgId string `json:"msg_id"`
}

func (b BaseRequest) HasError(w http.ResponseWriter) bool {
	err := []error{}
	helper.CheckEmpty(b.MsgId, &err, "msg_id")
	return len(err) != 0
}
