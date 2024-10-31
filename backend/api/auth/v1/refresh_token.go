package v1

import "github.com/gogf/gf/v2/frame/g"

type RefreshTokenReq struct {
	g.Meta `path:"/token" tags:"Auth" method:"get" sum:"Refresh user token"`
}
type RefreshTokenRes struct {
	Token  string `json:"token" des:"JWT token"`
	Expire int64  `json:"expire" des:"JWT expire time"`
}
