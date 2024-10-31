package users

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"overflows/api/users/v1"
)

func (c *ControllerV1) GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserReq) (res *v1.GetCurrentUserRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
