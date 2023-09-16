package response

type SuccessResponse struct {
	MsgId      string      `json:"msg_id"`
	StatusCode HttpStatus  `json:"status_code"`
	Response   interface{} `json:"response"`
}
