// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BetLogUserRebateDao is the data access object for the table bet_log_user_rebate.
type BetLogUserRebateDao struct {
	table    string                  // table is the underlying table name of the DAO.
	group    string                  // group is the database configuration group name of the current DAO.
	columns  BetLogUserRebateColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler      // handlers for customized model modification.
}

// BetLogUserRebateColumns defines and stores column names for the table bet_log_user_rebate.
type BetLogUserRebateColumns struct {
	Id                   string //
	SiteId               string // 站点ID
	RebateStartTime      string // 洗码开始时间
	RebateEndTime        string // 洗码结束时间
	GameId               string // 游戏ID
	AgentId              string // 代理ID
	GameType             string // 游戏类型
	UserId               string // 会员ID
	Username             string // 会员用户名
	BetAmount            string // 投注金额
	BetCount             string // 注单数量
	ValidBetAmount       string // 有效投注金额有效投注金额（投注时间）
	ValidBetAmountSettle string // 有效投注金额（结算时间）
	WinOrLose            string // 输赢结果
	CreatedAt            string //
	UpdatedAt            string //
}

// betLogUserRebateColumns holds the columns for the table bet_log_user_rebate.
var betLogUserRebateColumns = BetLogUserRebateColumns{
	Id:                   "id",
	SiteId:               "site_id",
	RebateStartTime:      "rebate_start_time",
	RebateEndTime:        "rebate_end_time",
	GameId:               "game_id",
	AgentId:              "agent_id",
	GameType:             "game_type",
	UserId:               "user_id",
	Username:             "username",
	BetAmount:            "bet_amount",
	BetCount:             "bet_count",
	ValidBetAmount:       "valid_bet_amount",
	ValidBetAmountSettle: "valid_bet_amount_settle",
	WinOrLose:            "win_or_lose",
	CreatedAt:            "created_at",
	UpdatedAt:            "updated_at",
}

// NewBetLogUserRebateDao creates and returns a new DAO object for table data access.
func NewBetLogUserRebateDao(handlers ...gdb.ModelHandler) *BetLogUserRebateDao {
	return &BetLogUserRebateDao{
		group:    "by_site",
		table:    "bet_log_user_rebate",
		columns:  betLogUserRebateColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BetLogUserRebateDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BetLogUserRebateDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BetLogUserRebateDao) Columns() BetLogUserRebateColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BetLogUserRebateDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BetLogUserRebateDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BetLogUserRebateDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
