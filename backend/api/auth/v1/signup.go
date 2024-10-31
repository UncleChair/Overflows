package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

// parameters should use uppercase
type SignupReq struct {
	g.Meta   `path:"/signup" tags:"Auth" method:"post" sum:"User Signup"`
	Username string `json:"username" v:"required|length:1,10" des:"User Name" eg:"admin"`
	Email    string `json:"email" v:"required|email" des:"User Email" eg:"admin@gmail.com"`
	Password string `json:"password" v:"required|password2" des:"Password" eg:"adminPassword"`
}
type SignupRes struct {
	Token string ` json:"token" des:"JWT token" eg:"TOOOOOOKEN"`
}
