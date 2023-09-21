package response

type LoginResponse struct {
	Success          bool   `json:"success"`
	Token            string `json:"token"`
	IsTokenRefreshed bool   `json:"is_token_refreshed"`
}
