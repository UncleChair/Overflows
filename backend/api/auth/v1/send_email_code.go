package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type SendEmailCodeReq struct {
	g.Meta   `path:"/email_code" tags:"Auth" method:"post" sum:"Set temp verification code to user"`
	Username string `json:"username" v:"required-without:Email" des:"User Name" eg:"admin"`
	Email    string `json:"email" v:"required-without:Username" des:"User Email" eg:"admin@gmail.com"`
}

type SendEmailCodeRes struct{}
