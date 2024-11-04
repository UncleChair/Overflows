package auth

import (
	"context"

	v1 "overflows/api/auth/v1"
	"overflows/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	res = &v1.LoginRes{}
	token, expire := service.JWTAuth().LoginHandler(ctx)
	res.Token = token
	res.Expire = expire.UnixMilli()
	return
}
