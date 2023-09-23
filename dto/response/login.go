package response

type LoginResponse struct {
	Success          bool   `json:"success"`
}

type RegisterResponse struct{
	Username string `json:"username"` 
}