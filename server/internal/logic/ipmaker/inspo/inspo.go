package inspo

import (
	"context"
	"fmt"

	"ip-maker-server/api/api/ipmaker/v1/inspo"
	"ip-maker-server/internal/model/input/appin"
	"ip-maker-server/internal/service/ipmaker"
)

type sInspo struct{}

func init() {
	ipmaker.RegisterInspo(&sInspo{})
}

func (s *sInspo) Create(ctx context.Context, in *appin.InspoCreateInput) (int64, error) {
	// TODO: dao.Inspo.Ctx(ctx).Data(g.Map{"title": in.Title, "type": in.Type, "content": in.Content, "operator": in.Operator}).InsertAndGetId()
	fmt.Println("Create inspo", in)
	return 1, nil
}

func (s *sInspo) Delete(ctx context.Context, id int64, operator int64) error {
	// TODO: dao.Inspo.Ctx(ctx).Where("id", id).Delete()
	fmt.Println("Delete inspo", id)
	return nil
}

func (s *sInspo) Update(ctx context.Context, in *appin.InspoUpdateInput) error {
	// TODO: dao.Inspo.Ctx(ctx).Where("id", in.Id).Data(in).Update()
	fmt.Println("Update inspo", in)
	return nil
}

func (s *sInspo) Get(ctx context.Context, id int64) (*inspo.InspoGetRes, error) {
	// TODO: dao.Inspo.Ctx(ctx).Where("id", id).Scan(&res)
	fmt.Println("Get inspo", id)
	return &inspo.InspoGetRes{}, nil
}

func (s *sInspo) List(ctx context.Context, in *appin.InspoListInput) ([]*inspo.InspoGetRes, int, error) {
	// TODO: m := dao.Inspo.Ctx(ctx); if in.Status != nil { m = m.Where("status", *in.Status) }
	// total, _ = m.Count(); m.Page(in.Page, in.Size).OrderDesc("created_at").Scan(&list)
	fmt.Println("List inspo", in)
	return []*inspo.InspoGetRes{}, 0, nil
}
