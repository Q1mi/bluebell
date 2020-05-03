package controller

const (
	CodeSuccess         = 1000
	CodeInvalidParams   = 1001
	CodeUserExist       = 1002
	CodeUserNotExist    = 1003
	CodeInvalidPassword = 1004
	CodeServerBusy      = 1005
)

var MsgFlags = map[int]string{
	CodeSuccess: "success",
	//CodeError:                      "fail",
	CodeInvalidParams:   "请求参数错误",
	CodeUserExist:       "用户名重复",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码错误",
	CodeServerBusy:      "服务繁忙",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[CodeServerBusy]
}
