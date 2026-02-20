// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// StoryDao is the data access object for the table hg_story.
type StoryDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  StoryColumns       // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// StoryColumns defines and stores column names for the table hg_story.
type StoryColumns struct {
	Id        string //
	IpId      string // 关联的IP ID
	Title     string // 故事标题
	Content   string // 故事内容
	CreatedAt string // 创建时间
	UpdatedAt string // 修改时间
}

// storyColumns holds the columns for the table hg_story.
var storyColumns = StoryColumns{
	Id:        "id",
	IpId:      "ip_id",
	Title:     "title",
	Content:   "content",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewStoryDao creates and returns a new DAO object for table data access.
func NewStoryDao(handlers ...gdb.ModelHandler) *StoryDao {
	return &StoryDao{
		group:    "default",
		table:    "hg_story",
		columns:  storyColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *StoryDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *StoryDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *StoryDao) Columns() StoryColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *StoryDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *StoryDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *StoryDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
