// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserDailyReport is the golang structure for table user_daily_report.
type UserDailyReport struct {
	Id                   uint64      `json:"id"                      orm:"id"                      description:""`             //
	SiteId               int         `json:"site_id"                 orm:"site_id"                 description:"站点ID"`         // 站点ID
	UserId               int         `json:"user_id"                 orm:"user_id"                 description:"会员ID"`         // 会员ID
	AgentId              int         `json:"agent_id"                orm:"agent_id"                description:"代理ID"`         // 代理ID
	Username             string      `json:"username"                orm:"username"                description:"用户名"`          // 用户名
	ReportDate           *gtime.Time `json:"report_date"             orm:"report_date"             description:"报表日期"`         // 报表日期
	BetCount             int         `json:"bet_count"               orm:"bet_count"               description:"注单数量"`         // 注单数量
	BetAmount            float64     `json:"bet_amount"              orm:"bet_amount"              description:"投注总额"`         // 投注总额
	ValidBetAmount       float64     `json:"valid_bet_amount"        orm:"valid_bet_amount"        description:"有效投注金额（投注时间）"` // 有效投注金额（投注时间）
	ValidBetAmountSettle float64     `json:"valid_bet_amount_settle" orm:"valid_bet_amount_settle" description:"有效投注金额（结算时间）"` // 有效投注金额（结算时间）
	WinOrLose            float64     `json:"win_or_lose"             orm:"win_or_lose"             description:"输赢结果"`         // 输赢结果
	RechargeAmount       float64     `json:"recharge_amount"         orm:"recharge_amount"         description:"充值金额"`         // 充值金额
	WithdrawAmount       float64     `json:"withdraw_amount"         orm:"withdraw_amount"         description:"提现金额"`         // 提现金额
	BonusAmount          float64     `json:"bonus_amount"            orm:"bonus_amount"            description:"红利金额"`         // 红利金额
	RebateAmount         float64     `json:"rebate_amount"           orm:"rebate_amount"           description:"返水金额"`         // 返水金额
	FeeAmount            float64     `json:"fee_amount"              orm:"fee_amount"              description:"手续费"`          // 手续费
	CreatedAt            *gtime.Time `json:"created_at"              orm:"created_at"              description:""`             //
	UpdatedAt            *gtime.Time `json:"updated_at"              orm:"updated_at"              description:""`             //
}
