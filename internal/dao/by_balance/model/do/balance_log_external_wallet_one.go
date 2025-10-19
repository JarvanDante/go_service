// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogExternalWalletOne is the golang structure of table balance_log_external_wallet_one for DAO operations like Where/Data.
type BalanceLogExternalWalletOne struct {
	g.Meta                `orm:"table:balance_log_external_wallet_one, do:true"`
	Id                    any         // ID
	SiteId                any         // 站点ID
	ChannelId             any         // 渠道id
	UserId                any         // 会员ID
	UserLevelId           any         // 会员层级ID
	Username              any         // 会员用户名
	TraceId               any         // 标识符UUID
	TransactionId         any         // 交易唯一ID
	BetId                 any         // 下注请求交易唯一ID
	ExternalTransactionId any         // 游戏供应商外部交易ID
	RoundId               any         // 游戏回合ID
	BetAmount             any         // 下注交易的金额
	WinAmount             any         // 赢取的金额
	EffectiveTurnover     any         // 有效投注额的金额
	WinLoss               any         // 绝对赢取或亏损的金额
	JackpotAmount         any         // 奖池金额，当jackpotAmount金额>0时需要钱包贷记操作
	ResultType            any         // '交易类型:"WIN"玩家赢得一笔下注"BET_WIN"玩家下注并赢得"BET_LOSE"玩家下注并输掉"LOSE"玩家输掉一笔下注"END"通知运营商回合已结束，不需要钱包借记或贷记操作
	IsFreespin            any         // 下注为免费旋转下注的状态
	IsEndRound            any         // 下注已完成的状态
	Currency              any         // 货币代码
	Token                 any         // OneAPI集成生成的用户会话令牌
	GameCode              any         // 游戏代码
	BetTime               any         // 此交易的初始请求的Unix时间戳（毫秒）
	SettledTime           any         // 下注结算的Unix时间戳（毫秒）此交易
	CreatedAt             *gtime.Time // created_at
	UpdatedAt             *gtime.Time // updated_at
}
