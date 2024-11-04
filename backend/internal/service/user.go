// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	auth "overflows/api/auth/v1"
	"overflows/internal/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
)

type (
	IUser interface {
		GenerateUid(ctx context.Context) (uid string)
		CreateUser(ctx context.Context, data *auth.SignupReq) (user *entity.Users, err error)
		// This function is used to find user from ctx after passing authenticator middleware
		GetUserFromCtx(ctx context.Context) *gdb.Model
		GetUserFromUid(ctx context.Context, uid string) (user *entity.Users, err error)
		// Users checked in this function are injected by jwt middleware in context
		CanManageProject(ctx context.Context, projectId int) (bool, error)
		// Users checked in this function are injected by jwt middleware in context
		CanViewProject(ctx context.Context, projectId int) (bool, error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
