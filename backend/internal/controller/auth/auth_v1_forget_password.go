package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"overflows/api/auth/v1"
)

func (c *ControllerV1) ForgetPassword(ctx context.Context, req *v1.ForgetPasswordReq) (res *v1.ForgetPasswordRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
