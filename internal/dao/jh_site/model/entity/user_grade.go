// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserGrade is the golang structure for table user_grade.
type UserGrade struct {
	Id                   uint        `json:"id"                     orm:"id"                     description:""`              //
	SiteId               int         `json:"site_id"                orm:"site_id"                description:"站点ID"`          // 站点ID
	Name                 string      `json:"name"                   orm:"name"                   description:"会员等级名称"`        // 会员等级名称
	PointsUpgrade        int         `json:"points_upgrade"         orm:"points_upgrade"         description:"升级消耗积分"`        // 升级消耗积分
	BonusUpgrade         float64     `json:"bonus_upgrade"          orm:"bonus_upgrade"          description:"额外返水比例: 体育"`    // 额外返水比例: 体育
	BonusBirthday        float64     `json:"bonus_birthday"         orm:"bonus_birthday"         description:"生日彩金"`          // 生日彩金
	RebatePercentSports  float64     `json:"rebate_percent_sports"  orm:"rebate_percent_sports"  description:"额外返水比例: 体育"`    // 额外返水比例: 体育
	RebatePercentLottery float64     `json:"rebate_percent_lottery" orm:"rebate_percent_lottery" description:"额外返水比例: 彩票"`    // 额外返水比例: 彩票
	RebatePercentLive    float64     `json:"rebate_percent_live"    orm:"rebate_percent_live"    description:"额外返水比例: 真人视讯"`  // 额外返水比例: 真人视讯
	RebatePercentEgame   float64     `json:"rebate_percent_egame"   orm:"rebate_percent_egame"   description:"额外返水比例: 电子游戏"`  // 额外返水比例: 电子游戏
	FieldsDisable        string      `json:"fields_disable"         orm:"fields_disable"         description:"字段开关，用来关闭哪些字段"` // 字段开关，用来关闭哪些字段
	AutoProviding        string      `json:"auto_providing"         orm:"auto_providing"         description:"哪些字段的业务是自动发放的"` // 哪些字段的业务是自动发放的
	Status               int         `json:"status"                 orm:"status"                 description:"状态.1=可用；0=禁用"`  // 状态.1=可用；0=禁用
	CreatedAt            *gtime.Time `json:"created_at"             orm:"created_at"             description:""`              //
	UpdatedAt            *gtime.Time `json:"updated_at"             orm:"updated_at"             description:""`              //
}
