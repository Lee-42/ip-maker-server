// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Ip is the golang structure for table ip.
type Ip struct {
	Id          int64       `json:"id"          orm:"id"          description:"ID"`
	Name        string      `json:"name"        orm:"name"        description:"ip名称"`
	Description string      `json:"description" orm:"description" description:"ip描述"`
	Avatar      string      `json:"avatar"      orm:"avatar"      description:"ip头像"`
	CreatedAt   *gtime.Time `json:"createdAt"   orm:"created_at"  description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   orm:"updated_at"  description:"更新时间"`
}
