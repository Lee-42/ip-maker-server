package ipmaker

import (
	"context"
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/appin"
)

type (
	IIp interface {
		Delete(ctx context.Context, in *appin.IpDeleteInput) (err error)
		Edit(ctx context.Context, in *appin.IpUpdateInput) (err error)
		Add(ctx context.Context, in *appin.IpCreateInput) (err error)
		List(ctx context.Context, in *appin.IpListInput) (res *appin.IpListOutput, err error)
		View(ctx context.Context, in *appin.IpViewInput) (res *entity.Ip, err error)
	}
)

var (
	localIp IIp
)

func Ip() IIp {
	if localIp == nil {
		panic("implement not found for interface IIp, forgot register?")
	}
	return localIp
}

func RegisterIp(i IIp) {
	localIp = i
}
