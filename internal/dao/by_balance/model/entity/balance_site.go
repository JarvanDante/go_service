// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceSite is the golang structure for table balance_site.
type BalanceSite struct {
	Id               uint        `json:"id"               orm:"id"                description:""`             //
	SiteId           int         `json:"siteId"           orm:"site_id"           description:"站点ID"`         // 站点ID
	ChannelId        int         `json:"channelId"        orm:"channel_id"        description:"渠道id"`         // 渠道id
	BalanceDefault   float64     `json:"balanceDefault"   orm:"balance_default"   description:"站点默认额度。默认50W"` // 站点默认额度。默认50W
	BalanceAvailable float64     `json:"balanceAvailable" orm:"balance_available" description:"站点可用额度"`       // 站点可用额度
	Currency         string      `json:"currency"         orm:"currency"          description:""`             //
	CreatedAt        *gtime.Time `json:"createdAt"        orm:"created_at"        description:""`             //
	UpdatedAt        *gtime.Time `json:"updatedAt"        orm:"updated_at"        description:""`             //
}
