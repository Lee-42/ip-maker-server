package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/appin"
)

// IpCreateReq is the request for creating a new ip.
type IpCreateReq struct {
	g.Meta `path:"/ip/create" method:"post" tags:"IP前台管理" summary:"创建IP"`
	*appin.IpCreateInput
}

// IpUpdateReq is the request for updating an existing ip.
type IpUpdateReq struct {
	g.Meta `path:"/ip/update" method:"post" tags:"IP前台管理" summary:"更新IP"`
	*appin.IpUpdateInput
}

// IpDeleteReq is the request for deleting an existing ip.
type IpDeleteReq struct {
	g.Meta `path:"/ip/delete" method:"post" tags:"IP前台管理" summary:"删除IP"`
	*appin.IpDeleteInput
}

// IpViewReq is the request for viewing an existing ip.
type IpViewReq struct {
	g.Meta `path:"/ip/view" method:"get" tags:"IP前台管理" summary:"获取单个IP信息"`
	*appin.IpViewInput
}

// IpListReq is the request for listing existing ips.
type IpListReq struct {
	g.Meta `path:"/ip/list" method:"get" tags:"IP前台管理" summary:"获取IP列表"`
	*appin.IpListInput
}

// IpListRes is the response for listing existing ips.
type IpListRes struct {
	*appin.IpListOutput
}

// IpViewRes is the response for viewing an existing ip.
type IpViewRes struct {
	*entity.Ip
}
