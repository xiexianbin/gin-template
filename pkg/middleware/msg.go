package middleware

var MsgFlags = map[int]string{
	SUCCESS:                      "ok",
	ERROR:                        "fail",
	INVALID_PARAMS:               "请求参数错误",
	ERROR_AUTH_CHECK_JWT_FAIL:    "JWT鉴权失败",
	ERROR_AUTH_CHECK_JWT_TIMEOUT: "JWT已超时",
	ERROR_AUTH_JWT:               "JWT生成失败",
	ERROR_AUTH:                   "JWT错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
