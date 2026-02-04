// Package admin
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package admin

import (
	"context"
	"hotgo/api/admin/meme/user"
	"hotgo/internal/service/meme"
)

var (
	AppUser = cAppUser{}
)

type cAppUser struct{}

// List 用户列表
func (c *cAppUser) List(ctx context.Context, req *user.ListReq) (res *user.ListRes, err error) {
	list, totalCount, err := meme.AppUser().List(ctx, &req.UserListInp)
	if err != nil {
		return
	}

	res = new(user.ListRes)
	res.List = list
	res.PageRes.Pack(req, totalCount)
	return
}

// View 查看用户详情
func (c *cAppUser) View(ctx context.Context, req *user.ViewReq) (res *user.ViewRes, err error) {
	data, err := meme.AppUser().View(ctx, &req.UserViewInp)
	if err != nil {
		return
	}

	res = &user.ViewRes{
		AppUser: data,
	}
	return
}

// Edit 编辑用户
func (c *cAppUser) Edit(ctx context.Context, req *user.EditReq) (res *user.EditRes, err error) {
	err = meme.AppUser().Edit(ctx, &req.UserEditInp)
	return
}

// Delete 删除用户
func (c *cAppUser) Delete(ctx context.Context, req *user.DeleteReq) (res *user.DeleteRes, err error) {
	err = meme.AppUser().Delete(ctx, &req.UserDeleteInp)
	return
}

// Status 更新用户状态
func (c *cAppUser) Status(ctx context.Context, req *user.StatusReq) (res *user.StatusRes, err error) {
	err = meme.AppUser().UpdateStatus(ctx, &req.UserStatusInp)
	return
}
