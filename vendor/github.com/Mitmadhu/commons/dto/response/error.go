package response

type HttpStatus uint64

type ErrorResponse struct {
	MsgId      string     `json:"msg_id"`
	Message    string     `json:"message"`
	StatusCode HttpStatus `json:"status_code"`
	Success    bool       `json:"success"`
}
