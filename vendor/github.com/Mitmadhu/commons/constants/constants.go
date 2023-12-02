package constants

const (
	//endpoints
	MYSQLDB      = "mysqlDB"
	USER_SERVICE = "user_service"

	ReqPtr         = "reqPtr"
	AccessToken    = "accessToken"
	RefreshToken   = "refreshToken"
	JWTValidation  = "jwtValidation"
	NoneValidation = "none"

	// errors
	StatusInternalServerError = "internal server error"
	TokenExipired             = "token has expired"
	InternalServerError       = "internal server error"
	UserNotFound              = "user not found"
	InvalidJWTToken           = "invalid JWT token"
	InvalidClaim              = "invalid claim"
	UsernameExists            = "username exists"
	BadRequest                = "bad request"
)
