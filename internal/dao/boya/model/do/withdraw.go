// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Withdraw is the golang structure of table withdraw for DAO operations like Where/Data.
type Withdraw struct {
	g.Meta    `orm:"table:withdraw, do:true"`
	Id        any         //
	Name      any         //
	Code      any         //
	Status    any         //
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
