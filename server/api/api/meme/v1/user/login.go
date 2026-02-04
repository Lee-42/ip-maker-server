// Package v1
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package user

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// LoginReq 用户登录请求
type LoginReq struct {
	g.Meta   `path:"/user/login" method:"post" tags:"C端用户" summary:"用户登录"`
	Username string `json:"username" v:"required#请输入用户名"`
	Password string `json:"password" v:"required#请输入密码"`
}

// LoginRes 用户登录响应
type LoginRes struct {
	Token     string        `json:"token" dc:"访问令牌"`
	ExpiresAt *gtime.Time   `json:"expiresAt" dc:"过期时间"`
	UserInfo  LoginUserInfo `json:"userInfo" dc:"用户信息"`
}

// LoginUserInfo 登录用户信息
type LoginUserInfo struct {
	Id       int64  `json:"id" dc:"用户ID"`
	Username string `json:"username" dc:"用户名"`
	Nickname string `json:"nickname" dc:"昵称"`
	Avatar   string `json:"avatar" dc:"头像"`
	Mobile   string `json:"mobile" dc:"手机号"`
	Email    string `json:"email" dc:"邮箱"`
}
