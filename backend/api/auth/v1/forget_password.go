package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ForgetPasswordReq struct {
	g.Meta   `path:"/forget_password" tags:"Auth" method:"post" sum:"Send reset password email"`
	Username string `json:"username" v:"required-without:Email" des:"User Name" eg:"admin"`
	Email    string `json:"email" v:"required-without:Username" des:"User Email" eg:"admin@gmail.com"`
}

type ForgetPasswordRes struct{}
