// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MemberBonusRecord is the golang structure for table member_bonus_record.
type MemberBonusRecord struct {
	Id        uint64      `json:"id"        orm:"id"         description:"主键ID"`                                 // 主键ID
	SiteId    uint64      `json:"siteId"    orm:"site_id"    description:"商户ID"`                                 // 商户ID
	UserId    uint64      `json:"userId"    orm:"user_id"    description:"会员ID"`                                 // 会员ID
	BonusType int         `json:"bonusType" orm:"bonus_type" description:"礼金类型: 1=周礼金 2=月礼金 3=升级礼金"`             // 礼金类型: 1=周礼金 2=月礼金 3=升级礼金
	Season    string      `json:"season"    orm:"season"     description:"周期标识: 周/2025W36, 月/2025-08, 升级/VIP等级"` // 周期标识: 周/2025W36, 月/2025-08, 升级/VIP等级
	Amount    float64     `json:"amount"    orm:"amount"     description:"领取金额"`                                 // 领取金额
	ExpiredAt *gtime.Time `json:"expiredAt" orm:"expired_at" description:"过期时间，升级礼金永不过期可为NULL"`                  // 过期时间，升级礼金永不过期可为NULL
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"领取时间"`                                 // 领取时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`                                 // 更新时间
}
