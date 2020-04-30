package models

type User struct {
	UserName  string `json:"username"`
	Password  string `json:"password"`
	CreatedOn int    `json:"-"`
}

func (u *User) PasswordValid() bool {
	return false
}
