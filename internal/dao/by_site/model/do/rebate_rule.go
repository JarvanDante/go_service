// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// RebateRule is the golang structure of table rebate_rule for DAO operations like Where/Data.
type RebateRule struct {
	g.Meta    `orm:"table:rebate_rule, do:true"`
	Id        any         //
	SiteId    any         //
	Name      any         //
	Status    any         //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
