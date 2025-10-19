// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SignInRecordsDao is the data access object for the table sign_in_records.
type SignInRecordsDao struct {
	table    string               // table is the underlying table name of the DAO.
	group    string               // group is the database configuration group name of the current DAO.
	columns  SignInRecordsColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler   // handlers for customized model modification.
}

// SignInRecordsColumns defines and stores column names for the table sign_in_records.
type SignInRecordsColumns struct {
	Id           string // 主键ID
	SiteId       string // 网站id
	UserId       string // 用户ID
	SignInDate   string // 签到日期（每天最多签到一次）
	RewardAmount string // 签到获得的奖励金额
	CreatedAt    string // 签到时间
	UpdatedAt    string // updated_at
}

// signInRecordsColumns holds the columns for the table sign_in_records.
var signInRecordsColumns = SignInRecordsColumns{
	Id:           "id",
	SiteId:       "site_id",
	UserId:       "user_id",
	SignInDate:   "sign_in_date",
	RewardAmount: "reward_amount",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewSignInRecordsDao creates and returns a new DAO object for table data access.
func NewSignInRecordsDao(handlers ...gdb.ModelHandler) *SignInRecordsDao {
	return &SignInRecordsDao{
		group:    "by_site",
		table:    "sign_in_records",
		columns:  signInRecordsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SignInRecordsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SignInRecordsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SignInRecordsDao) Columns() SignInRecordsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SignInRecordsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SignInRecordsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SignInRecordsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
