// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserMessage is the golang structure for table user_message.
type UserMessage struct {
	Id        uint        `json:"id"         orm:"id"         description:""`                    //
	SiteId    int         `json:"site_id"    orm:"site_id"    description:"站点ID"`                // 站点ID
	UserId    int         `json:"user_id"    orm:"user_id"    description:"用户ID"`                // 用户ID
	MessageId int         `json:"message_id" orm:"message_id" description:"消息ID"`                // 消息ID
	Status    int         `json:"status"     orm:"status"     description:"状态。1=已读；0=未读"`        // 状态。1=已读；0=未读
	DeleteAt  uint        `json:"delete_at"  orm:"delete_at"  description:"是否删除。0=未删除；其他为删除时间戳"` // 是否删除。0=未删除；其他为删除时间戳
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`                    //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`                    //
}
