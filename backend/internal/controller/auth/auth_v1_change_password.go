package auth

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	v1 "overflows/api/auth/v1"
	"overflows/internal/dao"
	"overflows/internal/model/entity"
	"overflows/internal/service"
)

func (c *ControllerV1) ChangePassword(ctx context.Context, req *v1.ChangePasswordReq) (res *v1.ChangePasswordRes, err error) {
	key := req.Uid + "_password"
	userModel := dao.Users.Ctx(ctx).Where(g.Map{dao.Users.Columns().Uid: req.Uid})
	var user *entity.Users
	err = userModel.Scan(&user)
	if err != nil {
		service.Context().SetHttpStatus(ctx, 500)
		return nil, gerror.NewCode(gcode.New(1, "", nil), "Database error")
	}
	if user == nil {
		service.Context().SetHttpStatus(ctx, 403)
		return nil, gerror.NewCode(gcode.New(1, "", nil), "User not exist")
	}
	if req.OldPassword != "" {
		if !service.Bcrypt().Compare(user.Password, req.OldPassword) {
			service.Context().SetHttpStatus(ctx, 403)
			return nil, gerror.NewCode(gcode.New(2, "", nil), "Wrong Password")
		}
	} else if req.Token != "" {
		if success, _ := service.Token().Verify(ctx, key, req.Token); !success {
			service.Context().SetHttpStatus(ctx, 403)
			return nil, gerror.NewCode(gcode.New(3, "", nil), "Wrong Token")
		}
	} else {
		service.Context().SetHttpStatus(ctx, 403)
		return nil, gerror.NewCode(gcode.New(4, "", nil), "Not Valid")
	}
	newPassword, err := service.Bcrypt().Generate(req.NewPassword)
	if err != nil {
		service.Context().SetHttpStatus(ctx, 500)
		return nil, gerror.NewCode(gcode.New(2, "", nil), "Encryption error")
	}
	userModel.Data(g.Map{
		dao.Users.Columns().Password: newPassword,
	}).OmitEmpty().Update()

	go service.MailServer().SendPasswordChangedEmail(ctx, user)

	url := g.Cfg().MustGet(ctx, "frontendURL")
	g.RequestFromCtx(ctx).Response.RedirectTo(url.String())

	return
}
