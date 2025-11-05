// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BetLogDaily is the golang structure for table bet_log_daily.
type BetLogDaily struct {
	Id                   uint        `json:"id"                      orm:"id"                      description:""`                   //
	SiteId               uint        `json:"site_id"                 orm:"site_id"                 description:"站点ID"`               // 站点ID
	BetDate              *gtime.Time `json:"bet_date"                orm:"bet_date"                description:"投注日期"`               // 投注日期
	GameId               uint        `json:"game_id"                 orm:"game_id"                 description:"游戏ID"`               // 游戏ID
	AgentId              int         `json:"agent_id"                orm:"agent_id"                description:"代理ID"`               // 代理ID
	GameType             uint        `json:"game_type"               orm:"game_type"               description:"游戏类型"`               // 游戏类型
	UserId               uint        `json:"user_id"                 orm:"user_id"                 description:"会员ID"`               // 会员ID
	Username             string      `json:"username"                orm:"username"                description:"会员用户名"`              // 会员用户名
	BetAmount            float64     `json:"bet_amount"              orm:"bet_amount"              description:"投注金额"`               // 投注金额
	BetCount             uint        `json:"bet_count"               orm:"bet_count"               description:"注单数量"`               // 注单数量
	ValidBetAmount       float64     `json:"valid_bet_amount"        orm:"valid_bet_amount"        description:"有效投注金额有效投注金额（投注时间）"` // 有效投注金额有效投注金额（投注时间）
	ValidBetAmountSettle float64     `json:"valid_bet_amount_settle" orm:"valid_bet_amount_settle" description:"有效投注金额（结算时间）"`       // 有效投注金额（结算时间）
	WinOrLose            float64     `json:"win_or_lose"             orm:"win_or_lose"             description:"输赢结果"`               // 输赢结果
	CreatedAt            *gtime.Time `json:"created_at"              orm:"created_at"              description:""`                   //
	UpdatedAt            *gtime.Time `json:"updated_at"              orm:"updated_at"              description:""`                   //
}
