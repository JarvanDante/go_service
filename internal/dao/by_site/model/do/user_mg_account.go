// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserMgAccount is the golang structure of table user_mg_account for DAO operations like Where/Data.
type UserMgAccount struct {
	g.Meta      `orm:"table:user_mg_account, do:true"`
	UserId      any         // 用户id
	MgAccountId any         // mg平台账号id
	CreatedAt   *gtime.Time // 创建时间
	UpdatedAt   *gtime.Time // 更新时间
}
