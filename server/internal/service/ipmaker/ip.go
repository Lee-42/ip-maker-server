package ipmaker

import (
	"context"
	"hotgo/addons/ip/model/entity"
	"hotgo/addons/ip/model/input/ipin"
)

type IIp interface {
	Delete(ctx context.Context, in *ipin.IpDeleteInput) (err error)
	Edit(ctx context.Context, in *ipin.IpUpdateInput) (err error)
	Add(ctx context.Context, in *ipin.IpCreateInput) (err error)
	List(ctx context.Context, in *ipin.IpListInput) (res *ipin.IpListOutput, err error)
	View(ctx context.Context, in *ipin.IpViewInput) (res *entity.Ip, err error)
}
