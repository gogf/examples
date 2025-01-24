package v1

import "github.com/gogf/gf/v2/frame/g"

type CreateReq struct {
	g.Meta `path:"/user" tags:"用户服务" method:"post" summary:"创建用户"`
	Name   string `v:"required"`
}
type CreateRes struct {
	Id string `json:"id" dc:"用户ID"`
}

type GetOneReq struct {
	g.Meta `path:"/user/{id}" tags:"用户服务" method:"get" summary:"查询详情"`
	Id     string `v:"required" dc:"用户ID"`
}
type GetOneRes struct {
	Data ListItem `json:"data" dc:"用户信息"`
}

type GetListReq struct {
	g.Meta `path:"/user" tags:"用户服务" method:"get" summary:"查询列表"`
	Ids    []string `v:"required" dc:"用户ID列表"`
}
type GetListRes struct {
	List []ListItem `json:"list" dc:"用户列表"`
}

type ListItem struct {
	Id        string `json:"id"         dc:"用户ID"`
	Name      string `json:"name"       dc:"用户名"`
	CreatedAt int64  `json:"created_at" dc:"创建时间"`
}

type DeleteReq struct {
	g.Meta `path:"/user" tags:"用户服务" method:"delete" summary:"删除用户"`
	Id     string `v:"required"`
}
type DeleteRes struct{}
