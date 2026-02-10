package user

import (
	"context"

	"github.com/gogf/gf-demo-user/v2/api/user/v1"
	usersvc `github.com/gogf/gf-demo-user/v2/internal/service/user`
)

// SignUp signs up the user.
func (c *ControllerV1) SignUp(ctx context.Context, req *v1.SignUpReq) (res *v1.SignUpRes, err error) {
	err = c.userSvc.Create(ctx, usersvc.CreateInput{
		Passport: req.Passport,
		Password: req.Password,
		Nickname: req.Nickname,
	})
	return
}
