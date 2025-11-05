// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteSign is the golang structure for table site_sign.
type SiteSign struct {
	Id        uint        `json:"id"         orm:"id"         description:""`                    //
	SiteId    int         `json:"site_id"    orm:"site_id"    description:"站点ID"`                // 站点ID
	Name      string      `json:"name"       orm:"name"       description:"活动名称"`                // 活动名称
	StartTime *gtime.Time `json:"start_time" orm:"start_time" description:"开始时间"`                // 开始时间
	EndTime   *gtime.Time `json:"end_time"   orm:"end_time"   description:"结束时间"`                // 结束时间
	Status    int         `json:"status"     orm:"status"     description:"活动状态。1=开启；0=关闭"`      // 活动状态。1=开启；0=关闭
	UserGrade string      `json:"user_grade" orm:"user_grade" description:"会员等级。以,的形式隔开"`        // 会员等级。以,的形式隔开
	UserLevel string      `json:"user_level" orm:"user_level" description:"会员层级。以,的形式隔开"`        // 会员层级。以,的形式隔开
	Platform  int         `json:"platform"   orm:"platform"   description:"优惠终端。0=所有；1=网站；2=手机"` // 优惠终端。0=所有；1=网站；2=手机
	Remark    string      `json:"remark"     orm:"remark"     description:"活动描述"`                // 活动描述
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`                    //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`                    //
}
