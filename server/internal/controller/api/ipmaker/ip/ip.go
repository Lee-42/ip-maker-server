package ip

import (
	"context"
	"hotgo/api/api/ipmaker/v1/ip"
	"hotgo/internal/service/ipmaker"
)

// CIp is the controller for ip.
var CIp = cIp{}

type cIp struct{}

// Delete deletes a ip.
func (c *cIp) Delete(ctx context.Context, req *v1.IpDeleteReq) (res *v1.IpDeleteReq, err error) {
	err = ipmaker.Ip().Delete(ctx, req.IpDeleteInput)
	return
}

// Edit edits a ip.
func (c *cIp) Edit(ctx context.Context, req *v1.IpUpdateReq) (res *v1.IpUpdateReq, err error) {
	err = ipmaker.Ip().Edit(ctx, req.IpUpdateInput)
	return
}

// Add adds a new ip.
func (c *cIp) Add(ctx context.Context, req *v1.IpCreateReq) (res *v1.IpCreateReq, err error) {
	err = ipmaker.Ip().Add(ctx, req.IpCreateInput)
	return
}

// List returns a list of ips.
func (c *cIp) List(ctx context.Context, req *v1.IpListReq) (res *v1.IpListRes, err error) {
	list, err := ipmaker.Ip().List(ctx, req.IpListInput)
	if err != nil {
		return
	}
	res = &v1.IpListRes{
		IpListOutput: list,
	}
	return
}

// View returns a ip.
func (c *cIp) View(ctx context.Context, req *v1.IpViewReq) (res *v1.IpViewRes, err error) {
	view, err := ipmaker.Ip().View(ctx, req.IpViewInput)
	if err != nil {
		return
	}
	res = &v1.IpViewRes{
		Ip: view,
	}
	return
}
