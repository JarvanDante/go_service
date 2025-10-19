// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogExternalWalletOne is the golang structure for table balance_log_external_wallet_one.
type BalanceLogExternalWalletOne struct {
	Id                    int64       `json:"id"                    orm:"id"                    description:"ID"`                                                                                                       // ID
	SiteId                int         `json:"siteId"                orm:"site_id"               description:"站点ID"`                                                                                                     // 站点ID
	ChannelId             int         `json:"channelId"             orm:"channel_id"            description:"渠道id"`                                                                                                     // 渠道id
	UserId                int         `json:"userId"                orm:"user_id"               description:"会员ID"`                                                                                                     // 会员ID
	UserLevelId           int         `json:"userLevelId"           orm:"user_level_id"         description:"会员层级ID"`                                                                                                   // 会员层级ID
	Username              string      `json:"username"              orm:"username"              description:"会员用户名"`                                                                                                    // 会员用户名
	TraceId               string      `json:"traceId"               orm:"traceId"               description:"标识符UUID"`                                                                                                  // 标识符UUID
	TransactionId         string      `json:"transactionId"         orm:"transactionId"         description:"交易唯一ID"`                                                                                                   // 交易唯一ID
	BetId                 string      `json:"betId"                 orm:"betId"                 description:"下注请求交易唯一ID"`                                                                                               // 下注请求交易唯一ID
	ExternalTransactionId string      `json:"externalTransactionId" orm:"externalTransactionId" description:"游戏供应商外部交易ID"`                                                                                              // 游戏供应商外部交易ID
	RoundId               string      `json:"roundId"               orm:"roundId"               description:"游戏回合ID"`                                                                                                   // 游戏回合ID
	BetAmount             float64     `json:"betAmount"             orm:"betAmount"             description:"下注交易的金额"`                                                                                                  // 下注交易的金额
	WinAmount             float64     `json:"winAmount"             orm:"winAmount"             description:"赢取的金额"`                                                                                                    // 赢取的金额
	EffectiveTurnover     float64     `json:"effectiveTurnover"     orm:"effectiveTurnover"     description:"有效投注额的金额"`                                                                                                 // 有效投注额的金额
	WinLoss               float64     `json:"winLoss"               orm:"winLoss"               description:"绝对赢取或亏损的金额"`                                                                                               // 绝对赢取或亏损的金额
	JackpotAmount         float64     `json:"jackpotAmount"         orm:"jackpotAmount"         description:"奖池金额，当jackpotAmount金额>0时需要钱包贷记操作"`                                                                         // 奖池金额，当jackpotAmount金额>0时需要钱包贷记操作
	ResultType            string      `json:"resultType"            orm:"resultType"            description:"'交易类型:\"WIN\"玩家赢得一笔下注\"BET_WIN\"玩家下注并赢得\"BET_LOSE\"玩家下注并输掉\"LOSE\"玩家输掉一笔下注\"END\"通知运营商回合已结束，不需要钱包借记或贷记操作"` // '交易类型:"WIN"玩家赢得一笔下注"BET_WIN"玩家下注并赢得"BET_LOSE"玩家下注并输掉"LOSE"玩家输掉一笔下注"END"通知运营商回合已结束，不需要钱包借记或贷记操作
	IsFreespin            int         `json:"isFreespin"            orm:"isFreespin"            description:"下注为免费旋转下注的状态"`                                                                                             // 下注为免费旋转下注的状态
	IsEndRound            int         `json:"isEndRound"            orm:"isEndRound"            description:"下注已完成的状态"`                                                                                                 // 下注已完成的状态
	Currency              string      `json:"currency"              orm:"currency"              description:"货币代码"`                                                                                                     // 货币代码
	Token                 string      `json:"token"                 orm:"token"                 description:"OneAPI集成生成的用户会话令牌"`                                                                                        // OneAPI集成生成的用户会话令牌
	GameCode              string      `json:"gameCode"              orm:"gameCode"              description:"游戏代码"`                                                                                                     // 游戏代码
	BetTime               int64       `json:"betTime"               orm:"betTime"               description:"此交易的初始请求的Unix时间戳（毫秒）"`                                                                                     // 此交易的初始请求的Unix时间戳（毫秒）
	SettledTime           int64       `json:"settledTime"           orm:"settledTime"           description:"下注结算的Unix时间戳（毫秒）此交易"`                                                                                      // 下注结算的Unix时间戳（毫秒）此交易
	CreatedAt             *gtime.Time `json:"createdAt"             orm:"created_at"            description:"created_at"`                                                                                               // created_at
	UpdatedAt             *gtime.Time `json:"updatedAt"             orm:"updated_at"            description:"updated_at"`                                                                                               // updated_at
}
