// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SignInRewardsByDayDao is the data access object for the table sign_in_rewards_by_day.
type SignInRewardsByDayDao struct {
	table    string                    // table is the underlying table name of the DAO.
	group    string                    // group is the database configuration group name of the current DAO.
	columns  SignInRewardsByDayColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler        // handlers for customized model modification.
}

// SignInRewardsByDayColumns defines and stores column names for the table sign_in_rewards_by_day.
type SignInRewardsByDayColumns struct {
	Id           string // 主键ID
	SiteId       string // 站点id
	DayNumber    string // 连续签到的第几天（从1开始）
	RewardAmount string // 第N天签到的奖励金额
	Remarks      string // 备注说明（可选）
	PaymentLimit string // 存款要求
	BetLimit     string // 投注要求
	CreatedAt    string // 创建时间
	UpdatedAt    string // 更新时间
}

// signInRewardsByDayColumns holds the columns for the table sign_in_rewards_by_day.
var signInRewardsByDayColumns = SignInRewardsByDayColumns{
	Id:           "id",
	SiteId:       "site_id",
	DayNumber:    "day_number",
	RewardAmount: "reward_amount",
	Remarks:      "remarks",
	PaymentLimit: "payment_limit",
	BetLimit:     "bet_limit",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

// NewSignInRewardsByDayDao creates and returns a new DAO object for table data access.
func NewSignInRewardsByDayDao(handlers ...gdb.ModelHandler) *SignInRewardsByDayDao {
	return &SignInRewardsByDayDao{
		group:    "by_site",
		table:    "sign_in_rewards_by_day",
		columns:  signInRewardsByDayColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SignInRewardsByDayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SignInRewardsByDayDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SignInRewardsByDayDao) Columns() SignInRewardsByDayColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SignInRewardsByDayDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SignInRewardsByDayDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SignInRewardsByDayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
