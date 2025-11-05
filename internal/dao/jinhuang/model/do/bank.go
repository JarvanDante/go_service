// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Bank is the golang structure of table bank for DAO operations like Where/Data.
type Bank struct {
	g.Meta    `orm:"table:bank, do:true"`
	Id        any         //
	Type      any         //
	Name      any         //
	Code      any         //
	Status    any         //
	Sort      any         //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
