// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentMessageLog is the golang structure for table agent_message_log.
type AgentMessageLog struct {
	Id             uint        `json:"id"               orm:"id"               description:""`       //
	SiteId         int         `json:"site_id"          orm:"site_id"          description:"站点ID"`   // 站点ID
	AgentMessageId int         `json:"agent_message_id" orm:"agent_message_id" description:"代理消息id"` // 代理消息id
	AgentId        int         `json:"agent_id"         orm:"agent_id"         description:"代理id"`   // 代理id
	IsRead         int         `json:"is_read"          orm:"is_read"          description:"是否读过"`   // 是否读过
	CreatedAt      *gtime.Time `json:"created_at"       orm:"created_at"       description:""`       //
	UpdatedAt      *gtime.Time `json:"updated_at"       orm:"updated_at"       description:""`       //
}
