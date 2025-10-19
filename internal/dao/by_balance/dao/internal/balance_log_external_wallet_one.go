// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceLogExternalWalletOneDao is the data access object for the table balance_log_external_wallet_one.
type BalanceLogExternalWalletOneDao struct {
	table    string                             // table is the underlying table name of the DAO.
	group    string                             // group is the database configuration group name of the current DAO.
	columns  BalanceLogExternalWalletOneColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler                 // handlers for customized model modification.
}

// BalanceLogExternalWalletOneColumns defines and stores column names for the table balance_log_external_wallet_one.
type BalanceLogExternalWalletOneColumns struct {
	Id                    string // ID
	SiteId                string // 站点ID
	ChannelId             string // 渠道id
	UserId                string // 会员ID
	UserLevelId           string // 会员层级ID
	Username              string // 会员用户名
	TraceId               string // 标识符UUID
	TransactionId         string // 交易唯一ID
	BetId                 string // 下注请求交易唯一ID
	ExternalTransactionId string // 游戏供应商外部交易ID
	RoundId               string // 游戏回合ID
	BetAmount             string // 下注交易的金额
	WinAmount             string // 赢取的金额
	EffectiveTurnover     string // 有效投注额的金额
	WinLoss               string // 绝对赢取或亏损的金额
	JackpotAmount         string // 奖池金额，当jackpotAmount金额>0时需要钱包贷记操作
	ResultType            string // '交易类型:"WIN"玩家赢得一笔下注"BET_WIN"玩家下注并赢得"BET_LOSE"玩家下注并输掉"LOSE"玩家输掉一笔下注"END"通知运营商回合已结束，不需要钱包借记或贷记操作
	IsFreespin            string // 下注为免费旋转下注的状态
	IsEndRound            string // 下注已完成的状态
	Currency              string // 货币代码
	Token                 string // OneAPI集成生成的用户会话令牌
	GameCode              string // 游戏代码
	BetTime               string // 此交易的初始请求的Unix时间戳（毫秒）
	SettledTime           string // 下注结算的Unix时间戳（毫秒）此交易
	CreatedAt             string // created_at
	UpdatedAt             string // updated_at
}

// balanceLogExternalWalletOneColumns holds the columns for the table balance_log_external_wallet_one.
var balanceLogExternalWalletOneColumns = BalanceLogExternalWalletOneColumns{
	Id:                    "id",
	SiteId:                "site_id",
	ChannelId:             "channel_id",
	UserId:                "user_id",
	UserLevelId:           "user_level_id",
	Username:              "username",
	TraceId:               "traceId",
	TransactionId:         "transactionId",
	BetId:                 "betId",
	ExternalTransactionId: "externalTransactionId",
	RoundId:               "roundId",
	BetAmount:             "betAmount",
	WinAmount:             "winAmount",
	EffectiveTurnover:     "effectiveTurnover",
	WinLoss:               "winLoss",
	JackpotAmount:         "jackpotAmount",
	ResultType:            "resultType",
	IsFreespin:            "isFreespin",
	IsEndRound:            "isEndRound",
	Currency:              "currency",
	Token:                 "token",
	GameCode:              "gameCode",
	BetTime:               "betTime",
	SettledTime:           "settledTime",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
}

// NewBalanceLogExternalWalletOneDao creates and returns a new DAO object for table data access.
func NewBalanceLogExternalWalletOneDao(handlers ...gdb.ModelHandler) *BalanceLogExternalWalletOneDao {
	return &BalanceLogExternalWalletOneDao{
		group:    "by_balance",
		table:    "balance_log_external_wallet_one",
		columns:  balanceLogExternalWalletOneColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BalanceLogExternalWalletOneDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BalanceLogExternalWalletOneDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BalanceLogExternalWalletOneDao) Columns() BalanceLogExternalWalletOneColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BalanceLogExternalWalletOneDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BalanceLogExternalWalletOneDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *BalanceLogExternalWalletOneDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
