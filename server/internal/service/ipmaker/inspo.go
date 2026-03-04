package ipmaker

import (
	"context"
	"ip-maker-server/api/api/ipmaker/v1/inspo"
	"ip-maker-server/internal/model/input/appin"
)

type IInspo interface {
	Create(ctx context.Context, in *appin.InspoCreateInput) (id int64, err error)
	Delete(ctx context.Context, id int64, operator int64) (err error)
	Update(ctx context.Context, in *appin.InspoUpdateInput) (err error)
	Get(ctx context.Context, id int64) (res *inspo.InspoGetRes, err error)
	List(ctx context.Context, in *appin.InspoListInput) (list []*inspo.InspoGetRes, total int, err error)
}

var localInspo IInspo

func Inspo() IInspo {
	if localInspo == nil {
		panic("implement not found for interface IInspo, forgot register?")
	}
	return localInspo
}

func RegisterInspo(i IInspo) {
	localInspo = i
}
