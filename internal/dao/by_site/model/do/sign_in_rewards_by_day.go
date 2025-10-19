// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// SignInRewardsByDay is the golang structure of table sign_in_rewards_by_day for DAO operations like Where/Data.
type SignInRewardsByDay struct {
	g.Meta       `orm:"table:sign_in_rewards_by_day, do:true"`
	Id           any         // 主键ID
	SiteId       any         // 站点id
	DayNumber    any         // 连续签到的第几天（从1开始）
	RewardAmount any         // 第N天签到的奖励金额
	Remarks      any         // 备注说明（可选）
	PaymentLimit any         // 存款要求
	BetLimit     any         // 投注要求
	CreatedAt    *gtime.Time // 创建时间
	UpdatedAt    *gtime.Time // 更新时间
}
