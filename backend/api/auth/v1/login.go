package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// parameters should use uppercase
type LoginReq struct {
	g.Meta   `path:"/login" tags:"Auth" method:"post" sum:"User Login"`
	Username string `json:"username" v:"required-without:Email" des:"User Name" eg:"admin"`
	Email    string `json:"email" v:"required-without:Username" des:"User Email" eg:"admin@gmail.com"`
	Password string `json:"password" v:"required" des:"Password" eg:"adminPassword"`
}
type LoginRes struct {
	Token  string `json:"token" des:"JWT token"`
	Expire int64  `json:"expire" des:"JWT expire time"`
}
