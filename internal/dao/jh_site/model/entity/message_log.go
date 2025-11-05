// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// MessageLog is the golang structure for table message_log.
type MessageLog struct {
	Id          uint        `json:"id"           orm:"id"           description:""`                          //
	SiteId      int         `json:"site_id"      orm:"site_id"      description:"站点ID"`                      // 站点ID
	MessageId   int         `json:"message_id"   orm:"message_id"   description:"消息ID"`                      // 消息ID
	MessageType int         `json:"message_type" orm:"message_type" description:"发送类型。1=会员等级，2=会员层级，2=会员状态"` // 发送类型。1=会员等级，2=会员层级，2=会员状态
	UserGrade   int         `json:"user_grade"   orm:"user_grade"   description:"会员等级"`                      // 会员等级
	UserLevel   int         `json:"user_level"   orm:"user_level"   description:"会员层级"`                      // 会员层级
	UserStatus  int         `json:"user_status"  orm:"user_status"  description:"会员状态.0=离线；1=在线"`            // 会员状态.0=离线；1=在线
	CreatedAt   *gtime.Time `json:"created_at"   orm:"created_at"   description:""`                          //
	UpdatedAt   *gtime.Time `json:"updated_at"   orm:"updated_at"   description:""`                          //
}
