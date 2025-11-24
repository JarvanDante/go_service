// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id                uint        `json:"id"                  orm:"id"                  description:""`                      //
	SiteId            int         `json:"site_id"             orm:"site_id"             description:""`                      //
	GradeId           int         `json:"grade_id"            orm:"grade_id"            description:""`                      //
	LevelId           int         `json:"level_id"            orm:"level_id"            description:""`                      //
	AgentId           int         `json:"agent_id"            orm:"agent_id"            description:""`                      //
	Username          string      `json:"username"            orm:"username"            description:""`                      //
	Password          string      `json:"password"            orm:"password"            description:""`                      //
	PayPassword       string      `json:"pay_password"        orm:"pay_password"        description:""`                      //
	Status            int         `json:"status"              orm:"status"              description:""`                      //
	RegisterIp        string      `json:"register_ip"         orm:"register_ip"         description:""`                      //
	RegisterTime      *gtime.Time `json:"register_time"       orm:"register_time"       description:""`                      //
	RegisterUrl       string      `json:"register_url"        orm:"register_url"        description:"注册来源"`                  // 注册来源
	RegisterDevice    int         `json:"register_device"     orm:"register_device"     description:"1=电脑；2=手机；3=平板"`        // 1=电脑；2=手机；3=平板
	LastLoginIp       string      `json:"last_login_ip"       orm:"last_login_ip"       description:""`                      //
	LastLoginTime     *gtime.Time `json:"last_login_time"     orm:"last_login_time"     description:""`                      //
	LastLoginAddress  string      `json:"last_login_address"  orm:"last_login_address"  description:"最后登录地址"`                // 最后登录地址
	Realname          string      `json:"realname"            orm:"realname"            description:""`                      //
	Mobile            string      `json:"mobile"              orm:"mobile"              description:""`                      //
	Email             string      `json:"email"               orm:"email"               description:"邮箱"`                    // 邮箱
	Qq                string      `json:"qq"                  orm:"qq"                  description:"QQ号"`                   // QQ号
	Birthday          *gtime.Time `json:"birthday"            orm:"birthday"            description:""`                      //
	Sex               int         `json:"sex"                 orm:"sex"                 description:"性别。0=未知；1=男；2=女"`       // 性别。0=未知；1=男；2=女
	IsOnline          int         `json:"is_online"           orm:"is_online"           description:""`                      //
	FocusLevel        int         `json:"focus_level"         orm:"focus_level"         description:"会员关注级别。1=正常；2=可疑；3=危险"` // 会员关注级别。1=正常；2=可疑；3=危险
	BalanceStatus     uint        `json:"balance_status"      orm:"balance_status"      description:"1=0="`                  // 1=0=
	SafeQuestion      string      `json:"safe_question"       orm:"safe_question"       description:"密保问题"`                  // 密保问题
	SafeAnswer        string      `json:"safe_answer"         orm:"safe_answer"         description:"密保答案"`                  // 密保答案
	ShowBeginnerGuide int         `json:"show_beginner_guide" orm:"show_beginner_guide" description:"是否显示新手引导。1=显示；0=不显示"`   // 是否显示新手引导。1=显示；0=不显示
	DeleteAt          uint        `json:"delete_at"           orm:"delete_at"           description:"是否删除。0=未删除；其他为删除时间戳"`   // 是否删除。0=未删除；其他为删除时间戳
	Remark            string      `json:"remark"              orm:"remark"              description:"备注"`                    // 备注
	CreatedAt         *gtime.Time `json:"created_at"          orm:"created_at"          description:""`                      //
	UpdatedAt         *gtime.Time `json:"updated_at"          orm:"updated_at"          description:""`                      //
	PayTimes          int         `json:"pay_times"           orm:"pay_times"           description:"充值次数"`                  // 充值次数
}
