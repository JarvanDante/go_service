// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RebateHistory is the golang structure for table rebate_history.
type RebateHistory struct {
	Id             uint        `json:"id"             orm:"id"               description:""`       //
	SiteId         uint        `json:"siteId"         orm:"site_id"          description:"站点ID"`   // 站点ID
	RebateDate     *gtime.Time `json:"rebateDate"     orm:"rebate_date"      description:"返水日期"`   // 返水日期
	UserCount      uint        `json:"userCount"      orm:"user_count"       description:"返水人数"`   // 返水人数
	ValidBetAmount float64     `json:"validBetAmount" orm:"valid_bet_amount" description:"有效投注金额"` // 有效投注金额
	Money          float64     `json:"money"          orm:"money"            description:"返水金额"`   // 返水金额
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"       description:""`       //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"       description:""`       //
}
