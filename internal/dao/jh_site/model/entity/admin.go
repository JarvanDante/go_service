// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Admin is the golang structure for table admin.
type Admin struct {
	Id                 uint        `json:"id"                   orm:"id"                   description:""`                               //
	SiteId             int         `json:"site_id"              orm:"site_id"              description:""`                               //
	Username           string      `json:"username"             orm:"username"             description:""`                               //
	Nickname           string      `json:"nickname"             orm:"nickname"             description:""`                               //
	Password           string      `json:"password"             orm:"password"             description:""`                               //
	AdminRoleId        int         `json:"admin_role_id"        orm:"admin_role_id"        description:""`                               //
	Status             int         `json:"status"               orm:"status"               description:""`                               //
	SwitchGoogle2Fa    int         `json:"switch_google2fa"     orm:"switch_google2fa"     description:"二次验证开关。1=打开；0=关闭"`               // 二次验证开关。1=打开；0=关闭
	Google2FaSecret    string      `json:"google2fa_secret"     orm:"google2fa_secret"     description:"二次验证密钥"`                         // 二次验证密钥
	TransferAuditSound int         `json:"transfer_audit_sound" orm:"transfer_audit_sound" description:"转账.审核提示声音控制。0=关闭 1=播放一次；2=循环播放"` // 转账.审核提示声音控制。0=关闭 1=播放一次；2=循环播放
	SoundLoopTime      string      `json:"sound_loop_time"      orm:"sound_loop_time"      description:"声音循环时间 单位秒"`                     // 声音循环时间 单位秒
	PaymentSound       int         `json:"payment_sound"        orm:"payment_sound"        description:"第三方支付提示声音控制。0=关闭 1=播放一次"`        // 第三方支付提示声音控制。0=关闭 1=播放一次
	LastLoginIp        string      `json:"last_login_ip"        orm:"last_login_ip"        description:""`                               //
	LastLoginTime      *gtime.Time `json:"last_login_time"      orm:"last_login_time"      description:""`                               //
	DeleteAt           uint        `json:"delete_at"            orm:"delete_at"            description:"是否删除。0=未删除；其他为删除时间戳"`            // 是否删除。0=未删除；其他为删除时间戳
	CreatedAt          *gtime.Time `json:"created_at"           orm:"created_at"           description:""`                               //
	UpdatedAt          *gtime.Time `json:"updated_at"           orm:"updated_at"           description:""`                               //
}
