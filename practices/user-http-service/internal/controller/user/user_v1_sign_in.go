package user

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/api/user/v1"
	usersvc `github.com/gogf/gf-demo-user/v2/internal/service/user`
)

// SignIn signs in the user.
func (c *ControllerV1) SignIn(ctx context.Context, req *v1.SignInReq) (res *v1.SignInRes, err error) {
	err = c.userSvc.SignIn(ctx, usersvc.SignInInput{
		Passport: req.Passport,
		Password: req.Password,
	})
	return
}
