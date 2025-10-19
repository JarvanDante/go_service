// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserMgAccountDao is the data access object for the table user_mg_account.
type UserMgAccountDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  UserMgAccountColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// UserMgAccountColumns defines and stores column names for the table user_mg_account.
type UserMgAccountColumns struct {
	UserId      string // 用户id
	MgAccountId string // mg平台账号id
	CreatedAt   string // 创建时间
	UpdatedAt   string // 更新时间
}

// userMgAccountColumns holds the columns for the table user_mg_account.
var userMgAccountColumns = UserMgAccountColumns{
	UserId:      "user_id",
	MgAccountId: "mg_account_id",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

// NewUserMgAccountDao creates and returns a new DAO object for table data access.
func NewUserMgAccountDao(handlers ...gdb.ModelHandler) *UserMgAccountDao {
	return &UserMgAccountDao{
		group:    "by_site",
		table:    "user_mg_account",
		columns:  userMgAccountColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserMgAccountDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserMgAccountDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserMgAccountDao) Columns() UserMgAccountColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserMgAccountDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserMgAccountDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserMgAccountDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
