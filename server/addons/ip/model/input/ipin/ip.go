package ipin

import (
	"hotgo/addons/ip/model/entity"
	"hotgo/internal/model/input/form"
)

// IpCreateInput is the input for creating a new ip.
type IpCreateInput struct {
	entity.Ip
}

// IpUpdateInput is the input for updating an existing ip.
type IpUpdateInput struct {
	entity.Ip
}

// IpDeleteInput is the input for deleting an existing ip.
type IpDeleteInput struct {
	Id int64 `json:"id" description:""`
}

// IpViewInput is the input for viewing an existing ip.
type IpViewInput struct {
	Id int64 `json:"id" description:""`
}

// IpListInput is the input for listing existing ips.
type IpListInput struct {
	form.PageReq
	Name string `json:"name" description:"ip名称"`
}

// IpListOutput is the output for listing existing ips.
type IpListOutput struct {
	List  []*entity.Ip `json:"list"  description:"列表"`
	Total int          `json:"total" description:"总数"`
}
