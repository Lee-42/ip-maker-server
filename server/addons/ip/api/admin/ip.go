package admin

import (
	"hotgo/addons/ip/model/entity"
	"hotgo/addons/ip/model/input/ipin"
)

// IpCreateReq is the request for creating a new ip.
type IpCreateReq struct {
	*ipin.IpCreateInput
}

// IpUpdateReq is the request for updating an existing ip.
type IpUpdateReq struct {
	*ipin.IpUpdateInput
}

// IpDeleteReq is the request for deleting an existing ip.
type IpDeleteReq struct {
	*ipin.IpDeleteInput
}

// IpViewReq is the request for viewing an existing ip.
type IpViewReq struct {
	*ipin.IpViewInput
}

// IpListReq is the request for listing existing ips.
type IpListReq struct {
	*ipin.IpListInput
}

// IpListRes is the response for listing existing ips.
type IpListRes struct {
	*ipin.IpListOutput
}

// IpViewRes is the response for viewing an existing ip.
type IpViewRes struct {
	*entity.Ip
}
