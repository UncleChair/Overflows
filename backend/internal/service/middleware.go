// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	IMiddleware interface {
		// Auth validates the request to allow only signed-in users visit.
		Auth(r *ghttp.Request)
		// CORS allows Cross-origin resource sharing.
		CORS(r *ghttp.Request)
		// i18n language setting.
		Language(r *ghttp.Request)
		// Prevent context canceled error.
		NeverDoneCtx(r *ghttp.Request)
		// Only admin could pass
		IsAdmin(r *ghttp.Request)
		// ResponseHandler is the middleware handling response object and its error.
		ResponseHandler(r *ghttp.Request)
		Ctx(r *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
