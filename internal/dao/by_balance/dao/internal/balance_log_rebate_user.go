// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceLogRebateUserDao is the data access object for the table balance_log_rebate_user.
type BalanceLogRebateUserDao struct {
	table    string                      // table is the underlying table name of the DAO.
	group    string                      // group is the database configuration group name of the current DAO.
	columns  BalanceLogRebateUserColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler          // handlers for customized model modification.
}

// BalanceLogRebateUserColumns defines and stores column names for the table balance_log_rebate_user.
type BalanceLogRebateUserColumns struct {
	SiteId          string // 站点ID
	ChannelId       string // 渠道id
	Id              string //
	UserId          string // 用户ID
	Username        string // 会员用户名
	TradeNo         string // 流水号
	RebateStartTime string // 洗码开始时间
	RebateEndTime   string // 洗码结束时间
	ValidBetAmount  string // 有效投注总金额
	Money           string // 返水金额
	Fee             string // 手续费
	Status          string // 状态。0=未确认；1=已确认
	AdminId         string // 后台管理员ID
	CreatedAt       string //
	UpdatedAt       string //
}

// balanceLogRebateUserColumns holds the columns for the table balance_log_rebate_user.
var balanceLogRebateUserColumns = BalanceLogRebateUserColumns{
	SiteId:          "site_id",
	ChannelId:       "channel_id",
	Id:              "id",
	UserId:          "user_id",
	Username:        "username",
	TradeNo:         "trade_no",
	RebateStartTime: "rebate_start_time",
	RebateEndTime:   "rebate_end_time",
	ValidBetAmount:  "valid_bet_amount",
	Money:           "money",
	Fee:             "fee",
	Status:          "status",
	AdminId:         "admin_id",
	CreatedAt:       "created_at",
	UpdatedAt:       "updated_at",
}

// NewBalanceLogRebateUserDao creates and returns a new DAO object for table data access.
func NewBalanceLogRebateUserDao(handlers ...gdb.ModelHandler) *BalanceLogRebateUserDao {
	return &BalanceLogRebateUserDao{
		group:    "by_balance",
		table:    "balance_log_rebate_user",
		columns:  balanceLogRebateUserColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BalanceLogRebateUserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BalanceLogRebateUserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BalanceLogRebateUserDao) Columns() BalanceLogRebateUserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BalanceLogRebateUserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BalanceLogRebateUserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BalanceLogRebateUserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
