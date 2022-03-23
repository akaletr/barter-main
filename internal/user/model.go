package user

import (
	"cmd/main/main.go/common"
)

type User struct {
	UserName    string       `json:"user_name"`
	Email       common.Email `json:"email"`
	Phone       common.Phone `json:"phone"`
	Description string       `json:"description"`
}
