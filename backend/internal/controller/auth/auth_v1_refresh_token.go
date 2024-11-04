package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "overflows/api/auth/v1"
	"overflows/internal/service"
)

func (c *ControllerV1) RefreshToken(ctx context.Context, req *v1.RefreshTokenReq) (res *v1.RefreshTokenRes, err error) {
	res = &v1.RefreshTokenRes{}
	token, expire, err := service.JWTAuth().RefreshToken(ctx)
	if err != nil {
		service.Context().SetHttpStatus(ctx, 401)
		err = gerror.NewCode(gcode.New(1, "Authentication failed", nil))
		res.Token = ""
		return
	}
	res.Token = token
	res.Expire = expire.UnixMilli()
	return
}
