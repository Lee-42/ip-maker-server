// Package ip
// @Link  https://github.com/bufanyun/hotgo
// @Copyright  Copyright (c) 2023 HotGo CLI
// @Author  Ms <133814250@qq.com>
// @License  https://github.com/bufanyun/hotgo/blob/master/LICENSE
package ip

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gctx"
	"hotgo/addons/ip/router"
	"hotgo/internal/library/addons"
	"hotgo/internal/service"
	"sync"
)

type module struct {
	skeleton *addons.Skeleton
	ctx      context.Context
	sync.Mutex
}

func init() {
	newModule()
}

func newModule() {
	m := &module{
		skeleton: &addons.Skeleton{
			Label:       "IP管理",
			Name:        "ip",
			Group:       1,
			Logo:        "",
			Brief:       "IP资源管理",
			Description: "提供IP资源的增删改查功能",
			Author:      "Your Name",
			Version:     "v1.0.0",
		},
		ctx: gctx.New(),
	}

	addons.RegisterModule(m)
}

// Start 启动模块
func (m *module) Start(option *addons.Option) (err error) {
	// 注册插件路由
	option.Server.Group("/", func(group *ghttp.RouterGroup) {
		group.Middleware(service.Middleware().Addon)
		router.Admin(m.ctx, group)
	})
	return
}

// Stop 停止模块
func (m *module) Stop() (err error) {
	return
}

// Ctx 上下文
func (m *module) Ctx() context.Context {
	return m.ctx
}

// GetSkeleton 获取模块
func (m *module) GetSkeleton() *addons.Skeleton {
	return m.skeleton
}

// Install 安装模块
func (m *module) Install(ctx context.Context) (err error) {
	// ...
	return
}

// Upgrade 更新模块
func (m *module) Upgrade(ctx context.Context) (err error) {
	// ...
	return
}

// UnInstall 卸载模块
func (m *module) UnInstall(ctx context.Context) (err error) {
	// ...
	return
}
