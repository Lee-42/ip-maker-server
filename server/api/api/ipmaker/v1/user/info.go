// Package v1
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// InfoReq 获取用户信息请求
type InfoReq struct {
	g.Meta `path:"/user/info" method:"get" tags:"C端用户" summary:"获取用户信息"`
}

// InfoRes 获取用户信息响应
type InfoRes struct {
	Id          int64       `json:"id" dc:"用户ID"`
	Username    string      `json:"username" dc:"用户名"`
	Nickname    string      `json:"nickname" dc:"昵称"`
	Avatar      string      `json:"avatar" dc:"头像"`
	Mobile      string      `json:"mobile" dc:"手机号"`
	Email       string      `json:"email" dc:"邮箱"`
	Sex         int         `json:"sex" dc:"性别"`
	Birthday    *gtime.Time `json:"birthday" dc:"生日"`
	LastLoginAt *gtime.Time `json:"lastLoginAt" dc:"最后登录时间"`
	LastLoginIp string      `json:"lastLoginIp" dc:"最后登录IP"`
	Status      int         `json:"status" dc:"状态"`
	CreatedAt   *gtime.Time `json:"createdAt" dc:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt" dc:"更新时间"`
}
