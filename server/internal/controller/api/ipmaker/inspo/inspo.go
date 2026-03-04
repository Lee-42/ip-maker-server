package inspo

import (
	"context"

	"ip-maker-server/api/api/ipmaker/v1/inspo"
	"ip-maker-server/internal/model/input/appin"
	"ip-maker-server/internal/service/ipmaker"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

func (c *Controller) Create(ctx context.Context, req *inspo.InspoCreateReq) (res *inspo.InspoCreateRes, err error) {
	in := &appin.InspoCreateInput{
		Title:    req.Title,
		Type:     req.Type,
		Content:  req.Content,
		Operator: 1, // 假定上下文提取, eg: contexts.GetUserId(ctx)
	}
	id, err := ipmaker.Inspo().Create(ctx, in)
	if err == nil {
		res = &inspo.InspoCreateRes{Id: id}
	}
	return
}

func (c *Controller) Delete(ctx context.Context, req *inspo.InspoDeleteReq) (res *inspo.InspoDeleteRes, err error) {
	operator := int64(1) // 假定上下文提取
	err = ipmaker.Inspo().Delete(ctx, req.Id, operator)
	return
}

func (c *Controller) Update(ctx context.Context, req *inspo.InspoUpdateReq) (res *inspo.InspoUpdateRes, err error) {
	in := &appin.InspoUpdateInput{
		Id:       req.Id,
		Title:    req.Title,
		Type:     req.Type,
		Content:  req.Content,
		Status:   req.Status,
		Operator: 1, // 假定上下文提取
	}
	err = ipmaker.Inspo().Update(ctx, in)
	return
}

func (c *Controller) Get(ctx context.Context, req *inspo.InspoGetReq) (res *inspo.InspoGetRes, err error) {
	res, err = ipmaker.Inspo().Get(ctx, req.Id)
	return
}

func (c *Controller) List(ctx context.Context, req *inspo.InspoListReq) (res *inspo.InspoListRes, err error) {
	in := &appin.InspoListInput{
		Page:   req.Page,
		Size:   req.Size,
		Status: req.Status,
	}
	list, total, err := ipmaker.Inspo().List(ctx, in)
	if err == nil {
		res = &inspo.InspoListRes{
			List:  list,
			Total: total,
		}
	}
	return
}
