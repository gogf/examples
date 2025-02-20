package user

import (
	"context"

	userSvcV1 "main/app/user/api/user/v1"

	"main/app/gateway/api/user/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	_, err = c.userSvc.Delete(ctx, &userSvcV1.DeleteReq{
		Id: req.Id,
	})
	return
}
