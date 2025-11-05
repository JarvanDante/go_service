// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceSite is the golang structure for table balance_site.
type BalanceSite struct {
	Id               uint        `json:"id"                orm:"id"                description:""`             //
	SiteId           int         `json:"site_id"           orm:"site_id"           description:"站点ID"`         // 站点ID
	BalanceDefault   float64     `json:"balance_default"   orm:"balance_default"   description:"站点默认额度。默认50W"` // 站点默认额度。默认50W
	BalanceAvailable float64     `json:"balance_available" orm:"balance_available" description:"站点可用额度"`       // 站点可用额度
	CreatedAt        *gtime.Time `json:"created_at"        orm:"created_at"        description:""`             //
	UpdatedAt        *gtime.Time `json:"updated_at"        orm:"updated_at"        description:""`             //
}
