package router

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"hotgo/addons/ip/controller/admin"
)

func Admin(ctx context.Context, group *ghttp.RouterGroup) {
	group.Group("/ip", func(group *ghttp.RouterGroup) {
		group.Bind(
			admin.CIp,
		)
	})
}
