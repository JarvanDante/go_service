// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BetLogUserRebate is the golang structure of table bet_log_user_rebate for DAO operations like Where/Data.
type BetLogUserRebate struct {
	g.Meta               `orm:"table:bet_log_user_rebate, do:true"`
	Id                   any         //
	SiteId               any         // 站点ID
	RebateStartTime      *gtime.Time // 洗码开始时间
	RebateEndTime        *gtime.Time // 洗码结束时间
	GameId               any         // 游戏ID
	AgentId              any         // 代理ID
	GameType             any         // 游戏类型
	UserId               any         // 会员ID
	Username             any         // 会员用户名
	BetAmount            any         // 投注金额
	BetCount             any         // 注单数量
	ValidBetAmount       any         // 有效投注金额有效投注金额（投注时间）
	ValidBetAmountSettle any         // 有效投注金额（结算时间）
	WinOrLose            any         // 输赢结果
	CreatedAt            *gtime.Time //
	UpdatedAt            *gtime.Time //
}
