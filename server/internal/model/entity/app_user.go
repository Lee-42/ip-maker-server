// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AppUser is the golang structure for table app_user.
type AppUser struct {
	Id           int64       `json:"id"           orm:"id"            description:"ID"`
	Username     string      `json:"username"     orm:"username"      description:"用户名"`
	Nickname     string      `json:"nickname"     orm:"nickname"      description:"昵称"`
	PasswordHash string      `json:"passwordHash" orm:"password_hash" description:"密码哈希"`
	Salt         string      `json:"salt"         orm:"salt"          description:"密码盐"`
	Mobile       string      `json:"mobile"       orm:"mobile"        description:"手机号"`
	Email        string      `json:"email"        orm:"email"         description:"邮箱"`
	Avatar       string      `json:"avatar"       orm:"avatar"        description:"头像URL"`
	Sex          int         `json:"sex"          orm:"sex"           description:"性别"`
	Birthday     *gtime.Time `json:"birthday"     orm:"birthday"      description:"生日"`
	LastLoginAt  *gtime.Time `json:"lastLoginAt"  orm:"last_login_at" description:"最后登录时间"`
	LastLoginIp  string      `json:"lastLoginIp"  orm:"last_login_ip" description:"最后登录IP"`
	Status       int         `json:"status"       orm:"status"        description:"状态"`
	Remark       string      `json:"remark"       orm:"remark"        description:"备注"`
	CreatedAt    *gtime.Time `json:"createdAt"    orm:"created_at"    description:"创建时间"`
	UpdatedAt    *gtime.Time `json:"updatedAt"    orm:"updated_at"    description:"更新时间"`
}
