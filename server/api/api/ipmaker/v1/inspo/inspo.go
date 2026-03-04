package inspo

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// --- 1. 创建 (Create) ---
type InspoCreateReq struct {
	g.Meta  `path:"/inspo" method:"post" tags:"Inspo前台管理" summary:"创建灵感"`
	Title   string `json:"title" v:"required#标题不能为空"`
	Type    string `json:"type" v:"required|in:text,image#请指定灵感类型(text/image)"`
	Content string `json:"content" v:"required#内容不能为空"`
}
type InspoCreateRes struct {
	Id int64 `json:"id"`
}

// --- 2. 删除 (Delete) ---
type InspoDeleteReq struct {
	g.Meta `path:"/inspo/{id}" method:"delete" tags:"Inspo前台管理" summary:"删除灵感"`
	Id     int64 `json:"id" v:"required#缺少灵感ID"`
}
type InspoDeleteRes struct{}

// --- 3. 更新 (Update) ---
type InspoUpdateReq struct {
	g.Meta  `path:"/inspo/{id}" method:"put" tags:"Inspo前台管理" summary:"更新灵感"`
	Id      int64   `json:"id" v:"required#缺少灵感ID"`
	Title   *string `json:"title"`
	Type    *string `json:"type" v:"in:text,image#灵感类型错误"`
	Content *string `json:"content"`
	Status  *int    `json:"status" v:"in:0,1#状态只能为0(未用)或1(已用)"`
}
type InspoUpdateRes struct{}

// --- 4. 详情查 (Get) ---
type InspoGetReq struct {
	g.Meta `path:"/inspo/{id}" method:"get" tags:"Inspo前台管理" summary:"获取灵感详情"`
	Id     int64 `json:"id" v:"required#缺少灵感ID"`
}
type InspoGetRes struct {
	Id        int64       `json:"id"`
	Title     string      `json:"title"`
	Type      string      `json:"type"`
	Content   string      `json:"content"`
	Status    int         `json:"status"` // 0:未使用 1:已使用
	CreatedAt *gtime.Time `json:"createdAt"`
	UpdatedAt *gtime.Time `json:"updatedAt"`
}

// --- 5. 列表查 (List) ---
type InspoListReq struct {
	g.Meta `path:"/inspo" method:"get" tags:"Inspo前台管理" summary:"获取灵感列表"`
	Page   int  `json:"page" d:"1"`
	Size   int  `json:"size" d:"20"`
	Status *int `json:"status" v:"in:0,1"`
}
type InspoListRes struct {
	List  []*InspoGetRes `json:"list"`
	Total int            `json:"total"`
}
