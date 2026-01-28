// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// AppUser is the golang structure of table hg_app_user for DAO operations like Where/Data.
type AppUser struct {
	g.Meta       `orm:"table:hg_app_user, do:true"`
	Id           any         // ID
	Username     any         // 用户名
	Nickname     any         // 昵称
	PasswordHash any         // 密码哈希
	Salt         any         // 密码盐
	Mobile       any         // 手机号
	Email        any         // 邮箱
	Avatar       any         // 头像URL
	Sex          any         // 性别
	Birthday     *gtime.Time // 生日
	LastLoginAt  *gtime.Time // 最后登录时间
	LastLoginIp  any         // 最后登录IP
	Status       any         // 状态
	Remark       any         // 备注
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
