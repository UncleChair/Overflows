package auth

import (
	"context"

	v1 "overflows/api/auth/v1"
	"overflows/internal/service"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	service.JWTAuth().LogoutHandler(ctx)
	return
}
