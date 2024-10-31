package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gmeta"
)

type User struct {
	gmeta.Meta `orm:"table:users"`
	Uid        string `json:"uid" eg:"1000000000"`
	AvatarUrl  string `json:"avatarUrl" eg:"default"`
	Username   string `json:"username" eg:"chair"`
	Email      string `json:"email" eg:"1@email.com"`
}

type GetCurrentUserReq struct {
	g.Meta `path:"/current" tags:"User" method:"get" sum:"Get current user info"`
}

type GetCurrentUserRes *User
