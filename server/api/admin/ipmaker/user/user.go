// Package user
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package user

import (
	"hotgo/internal/model/entity"
	"hotgo/internal/model/input/appin"
	"hotgo/internal/model/input/form"

	"github.com/gogf/gf/v2/frame/g"
)

// ListReq 用户列表
type ListReq struct {
	g.Meta `path:"/user/list" method:"get" tags:"C端用户管理" summary:"获取C端用户列表"`
	appin.UserListInp
}

type ListRes struct {
	List []*entity.AppUser `json:"list" dc:"数据列表"`
	form.PageRes
}

// ViewReq 查看用户详情
type ViewReq struct {
	g.Meta `path:"/user/view" method:"get" tags:"C端用户管理" summary:"获取用户详情"`
	appin.UserViewInp
}

type ViewRes struct {
	*entity.AppUser
}

// EditReq 编辑用户
type EditReq struct {
	g.Meta `path:"/user/edit" method:"post" tags:"C端用户管理" summary:"编辑用户"`
	appin.UserEditInp
}

type EditRes struct{}

// DeleteReq 删除用户
type DeleteReq struct {
	g.Meta `path:"/user/delete" method:"post" tags:"C端用户管理" summary:"删除用户"`
	appin.UserDeleteInp
}

type DeleteRes struct{}

// StatusReq 更新用户状态
type StatusReq struct {
	g.Meta `path:"/user/status" method:"post" tags:"C端用户管理" summary:"更新用户状态"`
	appin.UserStatusInp
}

type StatusRes struct{}
