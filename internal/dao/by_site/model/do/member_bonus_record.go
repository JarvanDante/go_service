// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberBonusRecord is the golang structure of table member_bonus_record for DAO operations like Where/Data.
type MemberBonusRecord struct {
	g.Meta    `orm:"table:member_bonus_record, do:true"`
	Id        any         // 主键ID
	SiteId    any         // 商户ID
	UserId    any         // 会员ID
	BonusType any         // 礼金类型: 1=周礼金 2=月礼金 3=升级礼金
	Season    any         // 周期标识: 周/2025W36, 月/2025-08, 升级/VIP等级
	Amount    any         // 领取金额
	ExpiredAt *gtime.Time // 过期时间，升级礼金永不过期可为NULL
	CreatedAt *gtime.Time // 领取时间
	UpdatedAt *gtime.Time // 更新时间
}
