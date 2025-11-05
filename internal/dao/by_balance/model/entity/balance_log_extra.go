// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogExtra is the golang structure for table balance_log_extra.
type BalanceLogExtra struct {
	Id        int         `json:"id"        orm:"id"         description:"ID"`      // ID
	SiteId    int         `json:"siteId"    orm:"site_id"    description:"应用id"`    // 应用id
	ChannelId int         `json:"channelId" orm:"channel_id" description:"渠道id"`    // 渠道id
	UserId    int         `json:"userId"    orm:"user_id"    description:"用户id"`    // 用户id
	Username  string      `json:"username"  orm:"username"   description:"用户名"`     // 用户名
	Money     float64     `json:"money"     orm:"money"      description:"金额"`      // 金额
	AdminId   int         `json:"adminId"   orm:"admin_id"   description:"后台管理员id"` // 后台管理员id
	Type      int         `json:"type"      orm:"type"       description:"类型1加2减"`  // 类型1加2减
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:"创建时间"`    // 创建时间
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:"更新时间"`    // 更新时间
}
