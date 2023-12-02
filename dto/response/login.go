package response

type LoginResponse struct {
	BaseResponse
	Success          bool   `json:"success"`
}

type RegisterResponse struct{
	BaseResponse
	Username string `json:"username"` 
}