package ip

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/appin"
	"hotgo/internal/service/ipmaker"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// SIp is the service for ip.
type SIp struct{}

func init() {
	ipmaker.RegisterIp(NewIp())
}

// NewIp returns a new SIp.
func NewIp() *SIp {
	return &SIp{}
}

// Delete deletes a ip.
func (s *SIp) Delete(ctx context.Context, in *appin.IpDeleteInput) (err error) {
	_, err = g.Model("ip").Ctx(ctx).Where("id", in.Id).Delete()
	return
}

// Edit edits a ip.
func (s *SIp) Edit(ctx context.Context, in *appin.IpUpdateInput) (err error) {
	if in.Id <= 0 {
		return gerror.New("ID is required")
	}
	_, err = g.Model("ip").Ctx(ctx).Where("id", in.Id).Data(in.Ip).Update()
	return
}

// Add adds a new ip.
func (s *SIp) Add(ctx context.Context, in *appin.IpCreateInput) (err error) {
	_, err = g.Model("ip").Ctx(ctx).Data(in.Ip).Insert()
	return
}

// List returns a list of ips.
func (s *SIp) List(ctx context.Context, in *appin.IpListInput) (res *appin.IpListOutput, err error) {
	res = new(appin.IpListOutput)
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
	err = mod.Page(in.PageReq.Page, in.PageReq.PerPage).Scan(&list)
	if err != nil {
		return
	}
	res.List = list
	return
}

// View returns a ip.
func (s *SIp) View(ctx context.Context, in *appin.IpViewInput) (res *entity.Ip, err error) {
	err = g.Model("ip").Ctx(ctx).Where("id", in.Id).Scan(&res)
	return
}
