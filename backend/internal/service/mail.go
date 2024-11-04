// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"overflows/internal/model/entity"

	"github.com/wneessen/go-mail"
)

type (
	IMailServer interface {
		NewClient() (client *mail.Client, err error)
		SendEmail(ctx context.Context, m *mail.Msg) (err error)
		SendPasswordChangedEmail(ctx context.Context, user *entity.Users) (bool, error)
		SendRegisterEmail(ctx context.Context, user *entity.Users) (bool, error)
		SendResetPasswordEmail(ctx context.Context, user *entity.Users) (bool, error)
		SendVerificationCodeEmail(ctx context.Context, user *entity.Users) (bool, error)
	}
)

var (
	localMailServer IMailServer
)

func MailServer() IMailServer {
	if localMailServer == nil {
		panic("implement not found for interface IMailServer, forgot register?")
	}
	return localMailServer
}

func RegisterMailServer(i IMailServer) {
	localMailServer = i
}
