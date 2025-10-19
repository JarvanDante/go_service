// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RebateHistory is the golang structure of table rebate_history for DAO operations like Where/Data.
type RebateHistory struct {
	g.Meta         `orm:"table:rebate_history, do:true"`
	Id             any         //
	SiteId         any         // 站点ID
	RebateDate     *gtime.Time // 返水日期
	UserCount      any         // 返水人数
	ValidBetAmount any         // 有效投注金额
	Money          any         // 返水金额
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
