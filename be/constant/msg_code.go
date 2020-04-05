package constant

// school-eco 错误码
const (
	// school-eco 服务公用状态码
	DEAL_SUCCESS   = 10000 //请求处理成功整个项目统一为1
	INVALID_PARAMS = 10001

	// web前后端的错误码
	WEB_SERVER_DEAL_ERROR          = 20000
	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)

// jwt
const (
	JwtSecret = "michael.liu"
)

// 错误码和错误信息映射表
var MsgFlags = map[int]string{
	// 请求处理成功
	DEAL_SUCCESS: "ok",

	// web 服务的错误码
	INVALID_PARAMS:                 "invalid params",
	WEB_SERVER_DEAL_ERROR:          "book seller server fail", // book-seller服务出错
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[WEB_SERVER_DEAL_ERROR]
}
