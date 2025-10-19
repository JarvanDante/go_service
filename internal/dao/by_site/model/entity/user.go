// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id                uint        `json:"id"                orm:"id"                  description:""`                      //
	SiteId            int         `json:"siteId"            orm:"site_id"             description:""`                      //
	GradeId           int         `json:"gradeId"           orm:"grade_id"            description:""`                      //
	LevelId           int         `json:"levelId"           orm:"level_id"            description:""`                      //
	AgentId           int         `json:"agentId"           orm:"agent_id"            description:""`                      //
	ChannelId         int         `json:"channelId"         orm:"channel_id"          description:"渠道id"`                  // 渠道id
	Username          string      `json:"username"          orm:"username"            description:""`                      //
	Password          string      `json:"password"          orm:"password"            description:""`                      //
	PayPassword       string      `json:"payPassword"       orm:"pay_password"        description:""`                      //
	Status            int         `json:"status"            orm:"status"              description:""`                      //
	RegisterIp        string      `json:"registerIp"        orm:"register_ip"         description:""`                      //
	RegisterTime      *gtime.Time `json:"registerTime"      orm:"register_time"       description:""`                      //
	RegisterUrl       string      `json:"registerUrl"       orm:"register_url"        description:"注册来源"`                  // 注册来源
	RegisterDevice    int         `json:"registerDevice"    orm:"register_device"     description:"1;电脑2; 手机3;平台"`         // 1;电脑2; 手机3;平台
	DeviceCode        string      `json:"deviceCode"        orm:"device_code"         description:"设备码"`                   // 设备码
	LastLoginIp       string      `json:"lastLoginIp"       orm:"last_login_ip"       description:""`                      //
	LastLoginTime     *gtime.Time `json:"lastLoginTime"     orm:"last_login_time"     description:""`                      //
	LastLoginAddress  string      `json:"lastLoginAddress"  orm:"last_login_address"  description:"最后登录地址"`                // 最后登录地址
	Realname          string      `json:"realname"          orm:"realname"            description:""`                      //
	Mobile            string      `json:"mobile"            orm:"mobile"              description:""`                      //
	Email             string      `json:"email"             orm:"email"               description:"邮箱"`                    // 邮箱
	Qq                string      `json:"qq"                orm:"qq"                  description:"QQ号"`                   // QQ号
	Birthday          *gtime.Time `json:"birthday"          orm:"birthday"            description:""`                      //
	Sex               int         `json:"sex"               orm:"sex"                 description:"性别。0=未知；1=男；2=女"`       // 性别。0=未知；1=男；2=女
	IsOnline          int         `json:"isOnline"          orm:"is_online"           description:""`                      //
	FocusLevel        int         `json:"focusLevel"        orm:"focus_level"         description:"会员关注级别。1=正常；2=可疑；3=危险"` // 会员关注级别。1=正常；2=可疑；3=危险
	BalanceStatus     uint        `json:"balanceStatus"     orm:"balance_status"      description:"1=0="`                  // 1=0=
	SafeQuestion      string      `json:"safeQuestion"      orm:"safe_question"       description:"密保问题"`                  // 密保问题
	SafeAnswer        string      `json:"safeAnswer"        orm:"safe_answer"         description:"密保答案"`                  // 密保答案
	ShowBeginnerGuide int         `json:"showBeginnerGuide" orm:"show_beginner_guide" description:"是否显示新手引导。1=显示；0=不显示"`   // 是否显示新手引导。1=显示；0=不显示
	DeleteAt          uint        `json:"deleteAt"          orm:"delete_at"           description:"是否删除。0=未删除；其他为删除时间戳"`   // 是否删除。0=未删除；其他为删除时间戳
	Remark            string      `json:"remark"            orm:"remark"              description:"备注"`                    // 备注
	SiteCode          string      `json:"siteCode"          orm:"site_code"           description:"应用码"`                   // 应用码
	ChannelCode       string      `json:"channelCode"       orm:"channel_code"        description:"渠道码"`                   // 渠道码
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"          description:""`                      //
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"          description:""`                      //
	PayTimes          int         `json:"payTimes"          orm:"pay_times"           description:"充值次数"`                  // 充值次数
	HasPass           int         `json:"hasPass"           orm:"has_pass"            description:"是否设置密码"`                // 是否设置密码
	Currency          string      `json:"currency"          orm:"currency"            description:"用户币种"`                  // 用户币种
	WithdrawTimes     int         `json:"withdrawTimes"     orm:"withdraw_times"      description:"提现次数"`                  // 提现次数
}
