// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Ip is the golang structure for table ip.
type Ip struct {
	Id          int64       `json:"id"          description:""`
	Name        string      `json:"name"        description:"ip名称"`
	Description string      `json:"description" description:"ip描述"`
	Avatar      string      `json:"avatar"      description:"ip头像"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:"创建时间"`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:"修改时间"`
}
