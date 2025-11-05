// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PointsChange is the golang structure for table points_change.
type PointsChange struct {
	Id           uint        `json:"id"            orm:"id"            description:""`                   //
	SiteId       int         `json:"site_id"       orm:"site_id"       description:"站点ID"`               // 站点ID
	ChangeType   int         `json:"change_type"   orm:"change_type"   description:"变化类型。1=增加；0=减少"`     // 变化类型。1=增加；0=减少
	TradeType    int         `json:"trade_type"    orm:"trade_type"    description:"类型。1=签到;2=充值;3=投注"`  // 类型。1=签到;2=充值;3=投注
	UserId       int         `json:"user_id"       orm:"user_id"       description:"会员ID"`               // 会员ID
	Username     string      `json:"username"      orm:"username"      description:"会员用户名"`              // 会员用户名
	PointsOld    int         `json:"points_old"    orm:"points_old"    description:"原积分"`                // 原积分
	PointsChange int         `json:"points_change" orm:"points_change" description:"积分变动"`               // 积分变动
	PointsNew    int         `json:"points_new"    orm:"points_new"    description:"现积分"`                // 现积分
	Status       int         `json:"status"        orm:"status"        description:"状态。1=待处理；2=成功；3=失败"` // 状态。1=待处理；2=成功；3=失败
	AdminId      int         `json:"admin_id"      orm:"admin_id"      description:"后台员工ID"`             // 后台员工ID
	Remark       string      `json:"remark"        orm:"remark"        description:""`                   //
	CreatedAt    *gtime.Time `json:"created_at"    orm:"created_at"    description:""`                   //
	UpdatedAt    *gtime.Time `json:"updated_at"    orm:"updated_at"    description:""`                   //
	TradeNo      string      `json:"trade_no"      orm:"trade_no"      description:""`                   //
}
