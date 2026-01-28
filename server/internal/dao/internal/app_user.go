// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AppUserDao is the data access object for the table hg_app_user.
type AppUserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AppUserColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AppUserColumns defines and stores column names for the table hg_app_user.
type AppUserColumns struct {
	Id           string // ID
	Username     string // 用户名
	Nickname     string // 昵称
	PasswordHash string // 密码哈希
	Salt         string // 密码盐
	Mobile       string // 手机号
	Email        string // 邮箱
	Avatar       string // 头像URL
	Sex          string // 性别
	Birthday     string // 生日
	LastLoginAt  string // 最后登录时间
	LastLoginIp  string // 最后登录IP
	Status       string // 状态
	Remark       string // 备注
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
}

// appUserColumns holds the columns for the table hg_app_user.
var appUserColumns = AppUserColumns{
	Id:           "id",
	Username:     "username",
	Nickname:     "nickname",
	PasswordHash: "password_hash",
	Salt:         "salt",
	Mobile:       "mobile",
	Email:        "email",
	Avatar:       "avatar",
	Sex:          "sex",
	Birthday:     "birthday",
	LastLoginAt:  "last_login_at",
	LastLoginIp:  "last_login_ip",
	Status:       "status",
	Remark:       "remark",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewAppUserDao creates and returns a new DAO object for table data access.
func NewAppUserDao(handlers ...gdb.ModelHandler) *AppUserDao {
	return &AppUserDao{
		group:    "default",
		table:    "hg_app_user",
		columns:  appUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AppUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AppUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AppUserDao) Columns() AppUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AppUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AppUserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AppUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
