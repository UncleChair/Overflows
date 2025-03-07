package auth

import (
	"context"

	v1 "overflows/api/auth/v1"
	"overflows/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) Signup(ctx context.Context, req *v1.SignupReq) (res *v1.SignupRes, err error) {
	user, err := service.User().CreateUser(ctx, req)
	if err != nil {
		code := gerror.Code(err)
		if code == gcode.New(1, "", nil) || code == gcode.New(2, "", nil) {
			return nil, err
		} else {
			return nil, gerror.WrapCode(gcode.New(2, "", nil), err, "User create failed")
		}
	}
	e := service.Casbin().DefaultEnforcer()
	e.AddGroupingPolicy(user.Uid, "user")
	e.SavePolicy()
	// go service.MailServer().SendRegisterEmail(ctx, user)
	res = &v1.SignupRes{}
	res.Token, _ = service.JWTAuth().LoginHandler(ctx)
	service.Context().SetHttpStatus(ctx, 201)
	return res, nil
}
