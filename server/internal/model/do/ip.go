// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Ip is the golang structure of table hg_ip for DAO operations like Where/Data.
type Ip struct {
	g.Meta      `orm:"table:hg_ip, do:true"`
	Id          any         // ID
	Name        any         // ip名称
	Description any         // ip描述
	Avatar      any         // ip头像
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
