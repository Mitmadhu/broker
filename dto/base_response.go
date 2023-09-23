package dto

type BaseResponse struct {
	MsgID          string `json:"msg_id"`
	Success        bool   `json:"success"`
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	IsTokenRefresh bool   `json:"is_token_refreshed"`
}
