// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package users

import (
	"context"

	"overflows/api/users/v1"
)

type IUsersV1 interface {
	GetCurrentUser(ctx context.Context, req *v1.GetCurrentUserReq) (res *v1.GetCurrentUserRes, err error)
	UploadUserAvatar(ctx context.Context, req *v1.UploadUserAvatarReq) (res *v1.UploadUserAvatarRes, err error)
}
