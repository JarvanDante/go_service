// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserSignInStatusDao is the data access object for the table user_sign_in_status.
type UserSignInStatusDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  UserSignInStatusColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// UserSignInStatusColumns defines and stores column names for the table user_sign_in_status.
type UserSignInStatusColumns struct {
	Id             string // 用户ID，主键，一对一对应用户表
	UserId         string // uid
	SiteId         string // 网站id
	LastSignInDate string // 最后一次签到日期
	ContinuousDays string // 连续签到天数（断签后重置）
	TotalDays      string // 累计签到总天数（不断签）
	CreatedAt      string // created_at
	UpdatedAt      string // 记录更新时间
}

// userSignInStatusColumns holds the columns for the table user_sign_in_status.
var userSignInStatusColumns = UserSignInStatusColumns{
	Id:             "id",
	UserId:         "user_id",
	SiteId:         "site_id",
	LastSignInDate: "last_sign_in_date",
	ContinuousDays: "continuous_days",
	TotalDays:      "total_days",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
}

// NewUserSignInStatusDao creates and returns a new DAO object for table data access.
func NewUserSignInStatusDao(handlers ...gdb.ModelHandler) *UserSignInStatusDao {
	return &UserSignInStatusDao{
		group:    "by_site",
		table:    "user_sign_in_status",
		columns:  userSignInStatusColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserSignInStatusDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserSignInStatusDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserSignInStatusDao) Columns() UserSignInStatusColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserSignInStatusDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserSignInStatusDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserSignInStatusDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
