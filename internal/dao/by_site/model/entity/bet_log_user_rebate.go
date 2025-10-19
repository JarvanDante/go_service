// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BetLogUserRebate is the golang structure for table bet_log_user_rebate.
type BetLogUserRebate struct {
	Id                   uint        `json:"id"                   orm:"id"                      description:""`                   //
	SiteId               uint        `json:"siteId"               orm:"site_id"                 description:"站点ID"`               // 站点ID
	RebateStartTime      *gtime.Time `json:"rebateStartTime"      orm:"rebate_start_time"       description:"洗码开始时间"`             // 洗码开始时间
	RebateEndTime        *gtime.Time `json:"rebateEndTime"        orm:"rebate_end_time"         description:"洗码结束时间"`             // 洗码结束时间
	GameId               uint        `json:"gameId"               orm:"game_id"                 description:"游戏ID"`               // 游戏ID
	AgentId              int         `json:"agentId"              orm:"agent_id"                description:"代理ID"`               // 代理ID
	GameType             uint        `json:"gameType"             orm:"game_type"               description:"游戏类型"`               // 游戏类型
	UserId               uint        `json:"userId"               orm:"user_id"                 description:"会员ID"`               // 会员ID
	Username             string      `json:"username"             orm:"username"                description:"会员用户名"`              // 会员用户名
	BetAmount            float64     `json:"betAmount"            orm:"bet_amount"              description:"投注金额"`               // 投注金额
	BetCount             uint        `json:"betCount"             orm:"bet_count"               description:"注单数量"`               // 注单数量
	ValidBetAmount       float64     `json:"validBetAmount"       orm:"valid_bet_amount"        description:"有效投注金额有效投注金额（投注时间）"` // 有效投注金额有效投注金额（投注时间）
	ValidBetAmountSettle float64     `json:"validBetAmountSettle" orm:"valid_bet_amount_settle" description:"有效投注金额（结算时间）"`       // 有效投注金额（结算时间）
	WinOrLose            float64     `json:"winOrLose"            orm:"win_or_lose"             description:"输赢结果"`               // 输赢结果
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"              description:""`                   //
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"              description:""`                   //
}
