package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ChangePasswordReq struct {
	g.Meta      `path:"/password" tags:"Auth" method:"post" sum:"Change password"`
	Uid         string `json:"uid" v:"required" des:"User unique id" in:"query"`
	Token       string `json:"token" v:"required-without:OldPassword" des:"Token" in:"query"`
	OldPassword string `json:"old_password" v:"required-without:Token|password2" des:"Old password"`
	NewPassword string `json:"new_password" v:"required|password2" des:"New password"`
}

type ChangePasswordRes struct{}
