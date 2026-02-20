// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Story is the golang structure of table hg_story for DAO operations like Where/Data.
type Story struct {
	g.Meta    `orm:"table:hg_story, do:true"`
	Id        any         //
	IpId      any         // 关联的IP ID
	Title     any         // 故事标题
	Content   any         // 故事内容
	CreatedAt *gtime.Time // 创建时间
	UpdatedAt *gtime.Time // 修改时间
}
