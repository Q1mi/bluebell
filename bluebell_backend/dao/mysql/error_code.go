package mysql

import "errors"

var (
	ErrorUserExit      = errors.New("用户已存在")
	ErrorUserNotExit   = errors.New("用户不已存在")
	ErrorPasswordWrong = errors.New("密码错误")
	ErrorGenIDFailed   = errors.New("创建用户ID失败")
)
