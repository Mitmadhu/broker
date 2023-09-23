package response


type SuccessResponse struct {
	BaseResponse
	Response   interface{} `json:"response"` 
}
