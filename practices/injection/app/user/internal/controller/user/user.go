package user

import (
	"context"

	"main/app/user/internal/service/user"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "main/app/user/api/user/v1"
)

type Controller struct {
	v1.UnimplementedUserServer
	userSvc *user.Service
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterUserServer(s.Server, &Controller{
		userSvc: user.New(),
	})
}

func (c *Controller) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	result, err := c.userSvc.Create(ctx, req.Name)
	if err != nil {
		return nil, gerror.Wrap(err, "create user failed")
	}
	res = &v1.CreateRes{
		Id: result,
	}
	return
}

func (c *Controller) GetOne(ctx context.Context, req *v1.GetOneReq) (res *v1.GetOneRes, err error) {
	result, err := c.userSvc.GetById(ctx, req.Id)
	if err != nil {
		return nil, gerror.Wrap(err, "get user failed")
	}
	res = &v1.GetOneRes{
		Data: result,
	}
	return
}

func (c *Controller) GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error) {
	result, err := c.userSvc.GetList(ctx, req.Ids)
	if err != nil {
		return nil, gerror.Wrap(err, "get user list failed")
	}
	res = &v1.GetListRes{
		List: result,
	}
	return
}

func (c *Controller) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	err = c.userSvc.DeleteById(ctx, req.Id)
	return
}
