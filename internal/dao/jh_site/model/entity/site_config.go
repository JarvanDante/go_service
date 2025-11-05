// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// SiteConfig is the golang structure for table site_config.
type SiteConfig struct {
	Id                   uint        `json:"id"                     orm:"id"                     description:""`                                       //
	SiteId               int         `json:"site_id"                orm:"site_id"                description:"站点ID"`                                   // 站点ID
	SwitchRegister       int         `json:"switch_register"        orm:"switch_register"        description:"是否开放注册。1=开放;0=关闭"`                       // 是否开放注册。1=开放;0=关闭
	SwitchGrade          int         `json:"switch_grade"           orm:"switch_grade"           description:"是否开放会员等级。1=开放;0=关闭"`                     // 是否开放会员等级。1=开放;0=关闭
	RegisterTimeInterval int         `json:"register_time_interval" orm:"register_time_interval" description:"同一IP重复注册。设定时间内，同一IP将无法进行多次注册。0或留空表示不限制"` // 同一IP重复注册。设定时间内，同一IP将无法进行多次注册。0或留空表示不限制
	DefaultGradeId       int         `json:"default_grade_id"       orm:"default_grade_id"       description:"默认用户等级ID"`                               // 默认用户等级ID
	DefaultLevelId       int         `json:"default_level_id"       orm:"default_level_id"       description:"默认用户层级ID"`                               // 默认用户层级ID
	DefaultAgentId       int         `json:"default_agent_id"       orm:"default_agent_id"       description:"默认代理ID"`                                 // 默认代理ID
	IsClose              int         `json:"is_close"               orm:"is_close"               description:"是否关站。1=是；0=否"`                           // 是否关站。1=是；0=否
	MobileLogo           string      `json:"mobile_logo"            orm:"mobile_logo"            description:"手机端logo"`                                // 手机端logo
	MinWithdraw          int         `json:"min_withdraw"           orm:"min_withdraw"           description:"单笔最低提现金额"`                               // 单笔最低提现金额
	MaxWithdraw          int         `json:"max_withdraw"           orm:"max_withdraw"           description:"单笔最高提现金额"`                               // 单笔最高提现金额
	CloseReason          string      `json:"close_reason"           orm:"close_reason"           description:"关站原因"`                                   // 关站原因
	UrlAgentPc           string      `json:"url_agent_pc"           orm:"url_agent_pc"           description:"前台代理系统链接地址"`                             // 前台代理系统链接地址
	UrlAgentRegister     string      `json:"url_agent_register"     orm:"url_agent_register"     description:"代理推广地址"`                                 // 代理推广地址
	UrlService           string      `json:"url_service"            orm:"url_service"            description:"客服链接"`                                   // 客服链接
	UrlMobile            string      `json:"url_mobile"             orm:"url_mobile"             description:"手机域名"`                                   // 手机域名
	CreatedAt            *gtime.Time `json:"created_at"             orm:"created_at"             description:""`                                       //
	UpdatedAt            *gtime.Time `json:"updated_at"             orm:"updated_at"             description:""`                                       //
	SwitchSign           string      `json:"switch_sign"            orm:"switch_sign"            description:""`                                       //
}
