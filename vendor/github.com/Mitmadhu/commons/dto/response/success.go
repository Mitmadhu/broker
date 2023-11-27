package response

type BaseResponse struct {
	MsgID      string `json:"msg_id"`
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
}

type SuccessResponse struct {
	BaseResponse
	Response interface{} `json:"response"`
}
