// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberDayReport is the golang structure of table member_day_report for DAO operations like Where/Data.
type MemberDayReport struct {
	g.Meta        `orm:"table:member_day_report, do:true"`
	Id            any         // 主键ID
	SiteId        any         // 商户ID
	UserId        any         // 会员ID
	Date          *gtime.Time // 统计日期
	DepositAmount any         // 当日充值总额
	BetAmount     any         // 当日投注总额
	CreatedAt     *gtime.Time // 记录创建时间
	UpdatedAt     *gtime.Time // 记录更新时间
}
