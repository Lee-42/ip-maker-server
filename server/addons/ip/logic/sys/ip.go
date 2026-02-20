package sys

import (
	"context"
	"hotgo/addons/ip/model/entity"
	"hotgo/addons/ip/model/input/ipin"
	"hotgo/internal/service"
	"hotgo/internal/service/ipmaker"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// SIp is the service for ip.
type SIp struct{}

func init() {
	service.RegisterIPMaker(&sIp{})
}

// NewIp returns a new SIp.
func NewIp() *SIp {
	return &SIp{}
}

// Delete deletes a ip.
func (s *SIp) Delete(ctx context.Context, in *ipin.IpDeleteInput) (err error) {
	_, err = g.Model("ip").Ctx(ctx).Where("id", in.Id).Delete()
	return
}

// Edit edits a ip.
func (s *SIp) Edit(ctx context.Context, in *ipin.IpUpdateInput) (err error) {
	if in.Id <= 0 {
		return gerror.New("ID is required")
	}
	_, err = g.Model("ip").Ctx(ctx).Where("id", in.Id).Data(in.Ip).Update()
	return
}

// Add adds a new ip.
func (s *SIp) Add(ctx context.Context, in *ipin.IpCreateInput) (err error) {
	_, err = g.Model("ip").Ctx(ctx).Data(in.Ip).Insert()
	return
}

// List returns a list of ips.
func (s *SIp) List(ctx context.Context, in *ipin.IpListInput) (res *ipin.IpListOutput, err error) {
	res = new(ipin.IpListOutput)
	mod := g.Model("ip").Ctx(ctx)
	if in.Name != "" {
		mod = mod.WhereLike("name", "%"+in.Name+"%")
	}
	res.Total, err = mod.Count()
	if err != nil {
		return
	}
	if res.Total == 0 {
		res.List = []*entity.Ip{}
		return
	}
	var list []*entity.Ip
	err = mod.Page(in.PageReq.Page, in.PageReq.Limit).Scan(&list)
	if err != nil {
		return
	}
	res.List = list
	return
}

// View returns a ip.
func (s *SIp) View(ctx context.Context, in *ipin.IpViewInput) (res *entity.Ip, err error) {
	err = g.Model("ip").Ctx(ctx).Where("id", in.Id).Scan(&res)
	return
}

type sIp struct{}

func (s *sIp) Upload() ipmaker.IUpload {
	return nil
}

func (s *sIp) User() ipmaker.IAppUser {
	return nil
}

func (s *sIp) Ip() ipmaker.IIp {
	return NewIp()
}
