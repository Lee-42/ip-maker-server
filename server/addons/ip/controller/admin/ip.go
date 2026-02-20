package admin

import (
	"context"
	"hotgo/addons/ip/api/admin"
	"hotgo/internal/service"
)

// CIp is the controller for ip.
var CIp = cIp{}

type cIp struct{}

// Delete deletes a ip.
func (c *cIp) Delete(ctx context.Context, req *admin.IpDeleteReq) (res *admin.IpDeleteReq, err error) {
	err = service.IpMaker().Ip().Delete(ctx, req.IpDeleteInput)
	return
}

// Edit edits a ip.
func (c *cIp) Edit(ctx context.Context, req *admin.IpUpdateReq) (res *admin.IpUpdateReq, err error) {
	err = service.IpMaker().Ip().Edit(ctx, req.IpUpdateInput)
	return
}

// Add adds a new ip.
func (c *cIp) Add(ctx context.Context, req *admin.IpCreateReq) (res *admin.IpCreateReq, err error) {
	err = service.IpMaker().Ip().Add(ctx, req.IpCreateInput)
	return
}

// List returns a list of ips.
func (c *cIp) List(ctx context.Context, req *admin.IpListReq) (res *admin.IpListRes, err error) {
	list, err := service.IpMaker().Ip().List(ctx, req.IpListInput)
	if err != nil {
		return
	}
	res = &admin.IpListRes{
		IpListOutput: list,
	}
	return
}

// View returns a ip.
func (c *cIp) View(ctx context.Context, req *admin.IpViewReq) (res *admin.IpViewRes, err error) {
	view, err := service.IpMaker().Ip().View(ctx, req.IpViewInput)
	if err != nil {
		return
	}
	res = &admin.IpViewRes{
		Ip: view,
	}
	return
}
