package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "服务异常",
	INVALID_PARAMS: "请求参数错误",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
