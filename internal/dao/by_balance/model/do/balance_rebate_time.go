// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceRebateTime is the golang structure of table balance_rebate_time for DAO operations like Where/Data.
type BalanceRebateTime struct {
	g.Meta         `orm:"table:balance_rebate_time, do:true"`
	Id             any         // ID
	SiteId         any         // $site_id
	ChannelId      any         // 渠道id
	UserId         any         // user_id
	RebateDate     *gtime.Time // 返水日期
	UserCount      any         //
	AdminUsername  any         // admin_username
	ValidBetAmount any         // 有效投注金额
	Money          any         // 金额
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	Remark         any         //
}
