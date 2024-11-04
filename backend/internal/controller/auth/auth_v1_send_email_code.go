package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gerror"

	v1 "overflows/api/auth/v1"
	"overflows/internal/dao"
	"overflows/internal/model/entity"
	"overflows/internal/service"
)

func (c *ControllerV1) SendEmailCode(ctx context.Context, req *v1.SendEmailCodeReq) (res *v1.SendEmailCodeRes, err error) {
	var key string
	var user *entity.Users
	if req.Email != "" {
		key = dao.Users.Columns().Email
		dao.Users.Ctx(ctx).Where(key, req.Email).Scan(&user)
	} else if req.Username != "" {
		key = dao.Users.Columns().Username
		dao.Users.Ctx(ctx).Where(key, req.Username).Scan(&user)
	}
	if user == nil {
		service.Context().SetHttpStatus(ctx, 404)
		return nil, gerror.New("User not found")
	}
	go service.MailServer().SendVerificationCodeEmail(ctx, user)
	return
}
