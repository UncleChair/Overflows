package middleware

import (
	"context"
	"overflows/internal/dao"
	"overflows/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
)

type sVerification struct{}

func init() {
	service.RegisterVerification(New())
}

func New() service.IVerification {
	return &sVerification{}
}

func (s *sVerification) CheckUsername(ctx context.Context, username string) error {
	n, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Username, username).Count()
	if err != nil {
		return gerror.Wrap(err, "Database error")
	}
	if n > 0 {
		return gerror.New("Username has been used")
	}
	return nil
}

func (s *sVerification) CheckUserEmail(ctx context.Context, email string) error {
	n, err := dao.Users.Ctx(ctx).Where(dao.Users.Columns().Email, email).Count()
	if err != nil {
		return gerror.Wrap(err, "Database error")
	}
	if n > 0 {
		return gerror.New("Email has been used")
	}
	return nil
}
