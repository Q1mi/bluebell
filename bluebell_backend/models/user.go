package models

type User struct {
	UserID   uint64 `json:"user_id" db:"user_id"`
	UserName string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

type RegisterForm struct {
	UserName        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (r RegisterForm) Validate() (ok bool, errMsg string) {
	if len(r.UserName) == 0 || len(r.Password) == 0 {
		errMsg = "用户名或密码格式有误"
		return
	}
	if r.Password != r.ConfirmPassword {
		errMsg = "两次密码不一致"
		return
	}
	return true, ""
}

func (u User) LoginValidate() (ok bool, errMsg string) {
	if len(u.UserName) == 0 || len(u.Password) == 0 {
		errMsg = "用户名或密码格式有误"
		return
	}
	return true, ""
}
