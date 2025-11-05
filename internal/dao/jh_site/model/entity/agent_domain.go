// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentDomain is the golang structure for table agent_domain.
type AgentDomain struct {
	Id        uint        `json:"id"         orm:"id"         description:""`             //
	SiteId    int         `json:"site_id"    orm:"site_id"    description:"站点ID"`         // 站点ID
	AgentId   int         `json:"agent_id"   orm:"agent_id"   description:"代理ID"`         // 代理ID
	Domain    string      `json:"domain"     orm:"domain"     description:"域名"`           // 域名
	Status    int         `json:"status"     orm:"status"     description:"状态。1=正常；0=禁用"` // 状态。1=正常；0=禁用
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`             //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`             //
}
