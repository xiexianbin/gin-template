package middleware

const (
	SUCCESS        = 200
	ERROR          = 500
	INVALID_PARAMS = 400

	ERROR_AUTH_CHECK_JWT_FAIL    = 40001
	ERROR_AUTH_CHECK_JWT_TIMEOUT = 40002
	ERROR_AUTH_JWT               = 40003
	ERROR_AUTH                   = 40004
)
