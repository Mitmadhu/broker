package response

type UserDetailsResponse struct {
	Username string `json:"username"`
	Age      uint16 `json:"age"`
	Address  string `json:"address"`
}
