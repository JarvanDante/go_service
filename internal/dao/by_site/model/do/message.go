// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure of table message for DAO operations like Where/Data.
type Message struct {
	g.Meta    `orm:"table:message, do:true"`
	Id        any         //
	SiteId    any         // 站点ID
	AdminId   any         // 员工ID
	Title     any         // 消息标题
	Content   any         // 消息内容
	Receiver  any         // 接收者
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
}
