package v1

import "github.com/gogf/gf/v2/frame/g"

// LogoutReq 退出登录请求
type LogoutReq struct {
	g.Meta `path:"/user/logout" method:"post" tags:"C端用户" summary:"退出登录"`
}

// LogoutRes 退出登录响应
type LogoutRes struct {
}
