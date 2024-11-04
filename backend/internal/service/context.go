// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"overflows/internal/model"

	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IContext interface {
		Init(r *ghttp.Request, customCtx *model.Context)
		// Use certain context to get custom context and modify it
		Use(ctx context.Context) *model.Context
		// SetHttpStatus set http status to context
		SetHttpStatus(ctx context.Context, status int)
		// SetUserUid set user uid to context custom field
		SetUserUid(ctx context.Context, uid string)
	}
)

var (
	localContext IContext
)

func Context() IContext {
	if localContext == nil {
		panic("implement not found for interface IContext, forgot register?")
	}
	return localContext
}

func RegisterContext(i IContext) {
	localContext = i
}
