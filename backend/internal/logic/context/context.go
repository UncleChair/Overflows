package context

import (
	"context"
	"overflows/internal/model"
	"overflows/internal/service"

	"github.com/gogf/gf/v2/net/ghttp"
)

type sContext struct{}

func init() {
	service.RegisterContext(New())
}

func New() *sContext {
	return &sContext{}
}

func (s *sContext) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(model.ContextKey, customCtx)
}

// Use certain context to get custom context and modify it
func (s *sContext) Use(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// SetHttpStatus set http status to context
func (s *sContext) SetHttpStatus(ctx context.Context, status int) {
	s.Use(ctx).HttpStatus = status
}

// SetUserUid set user uid to context custom field
func (s *sContext) SetUserUid(ctx context.Context, uid string) {
	s.Use(ctx).User.Uid = uid
}
