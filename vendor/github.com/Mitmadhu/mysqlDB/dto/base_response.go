package dto

type BaseResponse struct {
	MsgID      string `json:"msg_id"`
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
