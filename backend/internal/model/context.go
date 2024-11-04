package model

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

const (
	ContextKey = "ServiceContext"
)

type Context struct {
	HttpStatus int
	Session    *ghttp.Session
	User       *ContextUser
	Data       g.Map
}
type ContextUser struct {
	Uid string // User Unique ID
}
