// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SitePoints is the golang structure for table site_points.
type SitePoints struct {
	Id                   uint        `json:"id"                     orm:"id"                     description:""` //
	SiteId               int         `json:"site_id"                orm:"site_id"                description:""` //
	SwitchSitePoints     int         `json:"switch_site_points"     orm:"switch_site_points"     description:""` //
	SwitchRechargePoints int         `json:"switch_recharge_points" orm:"switch_recharge_points" description:""` //
	EachRechargeAmount   float64     `json:"each_recharge_amount"   orm:"each_recharge_amount"   description:""` //
	EachRechargePoints   int         `json:"each_recharge_points"   orm:"each_recharge_points"   description:""` //
	SwitchBettingPoints  int         `json:"switch_betting_points"  orm:"switch_betting_points"  description:""` //
	EachBettingAmount    float64     `json:"each_betting_amount"    orm:"each_betting_amount"    description:""` //
	EachBettingPoints    int         `json:"each_betting_points"    orm:"each_betting_points"    description:""` //
	MaxDailyPoints       int         `json:"max_daily_points"       orm:"max_daily_points"       description:""` //
	CreatedAt            *gtime.Time `json:"created_at"             orm:"created_at"             description:""` //
	UpdatedAt            *gtime.Time `json:"updated_at"             orm:"updated_at"             description:""` //
}
