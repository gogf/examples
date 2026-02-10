package user

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/api/user/v1"
)

// IsSignedIn checks if the user is signed in.
func (c *ControllerV1) IsSignedIn(ctx context.Context, req *v1.IsSignedInReq) (res *v1.IsSignedInRes, err error) {
	res = &v1.IsSignedInRes{
		OK: c.userSvc.IsSignedIn(ctx),
	}
	return
}
