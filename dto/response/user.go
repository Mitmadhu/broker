package response

type UserResponse struct {
	Username string `json:"username"`
	Age      int16  `json:"age"`
	Address  string `json:"address"`
	Email    string `json:"email"`
}
