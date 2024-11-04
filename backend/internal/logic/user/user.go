package user

import (
	"context"
	"image/color"
	"image/png"
	"os"
	"overflows/internal/consts"
	"overflows/internal/dao"
	"overflows/internal/model/entity"
	"overflows/internal/service"

	auth "overflows/api/auth/v1"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/issue9/identicon/v2"
)

type sUser struct{}

func init() {
	if !gfile.Exists(consts.AvatarPath) {
		if err := gfile.Mkdir(consts.AvatarPath); err != nil {
			g.Log().Fatal(gctx.New(), err)
		}
	}
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (s *sUser) GenerateUid(ctx context.Context) (uid string) {
	uid = grand.Digits(10)
	check, _ := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Uid, uid).Count()
	for check > 0 {
		uid = grand.Digits(10)
		check, _ = dao.Users.Ctx(ctx).Where(dao.Users.Columns().Uid, uid).Count()
	}
	return
}

func (s *sUser) CreateUser(ctx context.Context, data *auth.SignupReq) (user *entity.Users, err error) {
	r := g.RequestFromCtx(ctx)
	if err := service.Verification().CheckUserEmail(r.Context(), data.Email); err != nil {
		service.Context().SetHttpStatus(ctx, 409)
		return nil, gerror.WrapCode(gcode.New(1, "", nil), err, "")
	}
	if err := service.Verification().CheckUsername(r.Context(), data.Username); err != nil {
		service.Context().SetHttpStatus(ctx, 409)
		return nil, gerror.WrapCode(gcode.New(2, "", nil), err, "")
	}
	if err := gconv.Struct(data, &user); err != nil {
		service.Context().SetHttpStatus(ctx, 500)
		return nil, gerror.WrapCode(gcode.New(3, "", nil), err, "Conversion failed")
	}
	uid := s.GenerateUid(ctx)
	user.Uid = uid
	user.Password, err = service.Bcrypt().Generate(data.Password)
	if err != nil {
		service.Context().SetHttpStatus(ctx, 500)
		return nil, gerror.WrapCode(gcode.New(4, "", nil), err, "Password hashing failed")
	}
	// only store relative path here
	uri, err := s.generateDefaultAvatar(uid)
	user.AvatarUrl = uri
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		service.Context().SetHttpStatus(ctx, 500)
		return nil, gerror.WrapCode(gcode.New(5, "", nil), err, "Start database transaction failed")
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()
	_, err = dao.Users.Ctx(ctx).TX(tx).Data(user).OmitEmpty().Insert()
	if err != nil {
		service.Context().SetHttpStatus(ctx, 500)
		return nil, gerror.WrapCode(gcode.New(6, "", nil), err, "Database insert failed")
	}
	return
}

// This function is used to find user from ctx after passing authenticator middleware
func (s *sUser) GetUserFromCtx(ctx context.Context) *gdb.Model {
	key := dao.Users.Columns().Uid
	uid := service.Context().Use(ctx).User.Uid
	return dao.Users.Ctx(ctx).WithAll().Where(key, uid).Hook(dao.Users.AvatarPathHook())
}

func (s *sUser) GetUserFromUid(ctx context.Context, uid string) (user *entity.Users, err error) {
	key := dao.Users.Columns().Uid
	err = dao.Users.Ctx(ctx).Where(key, uid).Scan(&user)
	return
}

func (s *sUser) generateDefaultAvatar(uid string) (uri string, err error) {
	img := identicon.Make(
		identicon.Style2,
		64,
		color.NRGBA{255, 255, 255, 255},
		color.NRGBA{
			uint8(grand.Intn(256)),
			uint8(grand.Intn(256)),
			uint8(grand.Intn(256)),
			255,
		},
		[]byte(uid))
	uri = gfile.Join(consts.AvatarPath, uid+".png")
	if !gfile.Exists(consts.AvatarPath) {
		gfile.Create(consts.AvatarPath)
	}
	fi, err := os.Create(uri)
	png.Encode(fi, img)
	fi.Close()
	return
}

// Users checked in this function are injected by jwt middleware in context
func (s *sUser) CanManageProject(ctx context.Context, projectId int) (bool, error) {
	e := service.Casbin().DefaultEnforcer()
	uid := service.Context().Use(ctx).User.Uid
	isAdmin, _ := e.HasRoleForUser(uid, "admin")
	if isAdmin {
		return true, nil
	}

	// Check project owner
	// projectOwnerCount, err := dao.Projects.Ctx(ctx).Where(g.Map{
	// 	dao.Projects.Columns().Id:      projectId,
	// 	dao.Projects.Columns().OwnerId: uid,
	// }).Count()
	// if err != nil {
	// 	service.Context().SetHttpStatus(ctx, 500)
	// 	return false, gerror.WrapCode(gcode.New(1, "", nil), err, "Database error")
	// }
	// if projectOwnerCount > 0 {
	// 	return true, nil
	// }
	return false, nil
}

// Users checked in this function are injected by jwt middleware in context
func (s *sUser) CanViewProject(ctx context.Context, projectId int) (bool, error) {
	e := service.Casbin().DefaultEnforcer()
	uid := service.Context().Use(ctx).User.Uid
	isAdmin, _ := e.HasRoleForUser(uid, "admin")
	if isAdmin {
		return true, nil
	}

	return false, nil
}
