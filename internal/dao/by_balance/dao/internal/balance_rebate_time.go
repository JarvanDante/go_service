// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceRebateTimeDao is the data access object for the table balance_rebate_time.
type BalanceRebateTimeDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  BalanceRebateTimeColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// BalanceRebateTimeColumns defines and stores column names for the table balance_rebate_time.
type BalanceRebateTimeColumns struct {
	Id             string // ID
	SiteId         string // $site_id
	ChannelId      string // 渠道id
	UserId         string // user_id
	RebateDate     string // 返水日期
	UserCount      string //
	AdminUsername  string // admin_username
	ValidBetAmount string // 有效投注金额
	Money          string // 金额
	CreatedAt      string //
	UpdatedAt      string //
	Remark         string //
}

// balanceRebateTimeColumns holds the columns for the table balance_rebate_time.
var balanceRebateTimeColumns = BalanceRebateTimeColumns{
	Id:             "id",
	SiteId:         "site_id",
	ChannelId:      "channel_id",
	UserId:         "user_id",
	RebateDate:     "rebate_date",
	UserCount:      "user_count",
	AdminUsername:  "admin_username",
	ValidBetAmount: "valid_bet_amount",
	Money:          "money",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	Remark:         "remark",
}

// NewBalanceRebateTimeDao creates and returns a new DAO object for table data access.
func NewBalanceRebateTimeDao(handlers ...gdb.ModelHandler) *BalanceRebateTimeDao {
	return &BalanceRebateTimeDao{
		group:    "by_balance",
		table:    "balance_rebate_time",
		columns:  balanceRebateTimeColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BalanceRebateTimeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BalanceRebateTimeDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BalanceRebateTimeDao) Columns() BalanceRebateTimeColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BalanceRebateTimeDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BalanceRebateTimeDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BalanceRebateTimeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
