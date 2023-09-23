package dto

type BaseResponse struct {
	AccessToken    string `json:"access_token"`
	RefreshToken   string `json:"refresh_token"`
	IsTokenRefresh bool   `json:"is_token_refreshed"`
}
