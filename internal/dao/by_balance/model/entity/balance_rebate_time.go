// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceRebateTime is the golang structure for table balance_rebate_time.
type BalanceRebateTime struct {
	Id             int         `json:"id"             orm:"id"               description:"ID"`             // ID
	SiteId         int         `json:"siteId"         orm:"site_id"          description:"$site_id"`       // $site_id
	ChannelId      int         `json:"channelId"      orm:"channel_id"       description:"渠道id"`           // 渠道id
	UserId         int         `json:"userId"         orm:"user_id"          description:"user_id"`        // user_id
	RebateDate     *gtime.Time `json:"rebateDate"     orm:"rebate_date"      description:"返水日期"`           // 返水日期
	UserCount      int         `json:"userCount"      orm:"user_count"       description:""`               //
	AdminUsername  string      `json:"adminUsername"  orm:"admin_username"   description:"admin_username"` // admin_username
	ValidBetAmount float64     `json:"validBetAmount" orm:"valid_bet_amount" description:"有效投注金额"`         // 有效投注金额
	Money          float64     `json:"money"          orm:"money"            description:"金额"`             // 金额
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:""`               //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:""`               //
	Remark         string      `json:"remark"         orm:"remark"           description:""`               //
}
