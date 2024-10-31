package users

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"overflows/api/users/v1"
)

func (c *ControllerV1) UploadUserAvatar(ctx context.Context, req *v1.UploadUserAvatarReq) (res *v1.UploadUserAvatarRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
