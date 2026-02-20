// Package v1
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package v1

import "github.com/gogf/gf/v2/frame/g"

// UpdateReq 更新用户信息请求
type UpdateReq struct {
	g.Meta   `path:"/user/update" method:"post" tags:"C端用户" summary:"更新用户信息"`
	Nickname string `json:"nickname" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
	Sex      int    `json:"sex" dc:"性别"`
	Birthday string `json:"birthday" dc:"生日"`
}

// UpdateRes 更新用户信息响应
type UpdateRes struct{}
