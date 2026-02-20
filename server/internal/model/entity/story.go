// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Story is the golang structure for table story.
type Story struct {
	Id        int64       `json:"id"        orm:"id"         description:""`
	IpId      int64       `json:"ipId"      orm:"ip_id"      description:"关联的IP ID"`
	Title     string      `json:"title"     orm:"title"      description:"故事标题"`
	Content   string      `json:"content"   orm:"content"    description:"故事内容"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"修改时间"`
}
