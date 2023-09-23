package response

import "github.com/Mitmadhu/broker/dto"


type SuccessResponse struct {
	MsgId      string      `json:"msg_id"`
	StatusCode HttpStatus  `json:"status_code"`
	dto.BaseResponse
	Response   interface{} `json:"response"` 
}
