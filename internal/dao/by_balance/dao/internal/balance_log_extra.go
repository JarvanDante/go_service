// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceLogExtraDao is the data access object for the table balance_log_extra.
type BalanceLogExtraDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  BalanceLogExtraColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// BalanceLogExtraColumns defines and stores column names for the table balance_log_extra.
type BalanceLogExtraColumns struct {
	Id        string // ID
	SiteId    string // 应用id
	ChannelId string // 渠道id
	UserId    string // 用户id
	Username  string // 用户名
	Money     string // 金额
	AdminId   string // 后台管理员id
	Type      string // 类型1加2减
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// balanceLogExtraColumns holds the columns for the table balance_log_extra.
var balanceLogExtraColumns = BalanceLogExtraColumns{
	Id:        "id",
	SiteId:    "site_id",
	ChannelId: "channel_id",
	UserId:    "user_id",
	Username:  "username",
	Money:     "money",
	AdminId:   "admin_id",
	Type:      "type",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewBalanceLogExtraDao creates and returns a new DAO object for table data access.
func NewBalanceLogExtraDao(handlers ...gdb.ModelHandler) *BalanceLogExtraDao {
	return &BalanceLogExtraDao{
		group:    "by_balance",
		table:    "balance_log_extra",
		columns:  balanceLogExtraColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BalanceLogExtraDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BalanceLogExtraDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BalanceLogExtraDao) Columns() BalanceLogExtraColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BalanceLogExtraDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BalanceLogExtraDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BalanceLogExtraDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
