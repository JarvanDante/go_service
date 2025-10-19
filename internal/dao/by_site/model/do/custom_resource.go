// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CustomResource is the golang structure of table custom_resource for DAO operations like Where/Data.
type CustomResource struct {
	g.Meta         `orm:"table:custom_resource, do:true"`
	Id             any         //
	SiteId         any         // 站点id
	UserId         any         // 用户ID
	Ip             any         // ip
	ResourceDomain any         // 来源域名
	ResourceLink   any         // 原来链接
	Remark         any         // 备注
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
}
