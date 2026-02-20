// Package v1
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package v1

import "github.com/gogf/gf/v2/frame/g"

// RegisterReq 用户注册请求
type RegisterReq struct {
	g.Meta   `path:"/user/register" method:"post" tags:"C端用户" summary:"用户注册"`
	Username string `json:"username" v:"required|length:4,20#请输入用户名|用户名长度为4-20位"`
	Password string `json:"password" v:"required|length:6,20#请输入密码|密码长度为6-20位"`
	Mobile   string `json:"mobile" v:"phone#手机号格式不正确"`
	Email    string `json:"email" v:"email#邮箱格式不正确"`
}

// RegisterRes 用户注册响应
type RegisterRes struct {
	UserId int64 `json:"userId" dc:"用户ID"`
}
