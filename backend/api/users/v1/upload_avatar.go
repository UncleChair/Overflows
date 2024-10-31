package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type UploadUserAvatarReq struct {
	g.Meta `path:"/avatar" tags:"User" method:"post" sum:"Upload user avatar"`
	File   *ghttp.UploadFile `json:"file" type:"file" v:"required" des:"User Avatar"`
}

type UploadUserAvatarRes struct {
}
