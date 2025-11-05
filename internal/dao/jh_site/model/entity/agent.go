// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Agent is the golang structure for table agent.
type Agent struct {
	Id                  uint        `json:"id"                    orm:"id"                    description:""`                      //
	SiteId              int         `json:"site_id"               orm:"site_id"               description:"站点ID"`                  // 站点ID
	Username            string      `json:"username"              orm:"username"              description:"用户名"`                   // 用户名
	Password            string      `json:"password"              orm:"password"              description:"密码"`                    // 密码
	Status              int         `json:"status"                orm:"status"                description:"审核状态。0=停用；1=启用"`        // 审核状态。0=停用；1=启用
	StatusCheck         int         `json:"status_check"          orm:"status_check"          description:"审核状态。0=待审核；1=通过；2=未通过"` // 审核状态。0=待审核；1=通过；2=未通过
	Level               int         `json:"level"                 orm:"level"                 description:"代理层级"`                  // 代理层级
	Realname            string      `json:"realname"              orm:"realname"              description:"真实姓名"`                  // 真实姓名
	Mobile              string      `json:"mobile"                orm:"mobile"                description:"手机号"`                   // 手机号
	Email               string      `json:"email"                 orm:"email"                 description:"电子邮箱"`                  // 电子邮箱
	Qq                  string      `json:"qq"                    orm:"qq"                    description:"QQ"`                    // QQ
	Skype               string      `json:"skype"                 orm:"skype"                 description:"Skype"`                 // Skype
	BankName            string      `json:"bank_name"             orm:"bank_name"             description:"银行名称"`                  // 银行名称
	CardAccount         string      `json:"card_account"          orm:"card_account"          description:"银行户名"`                  // 银行户名
	CardNo              string      `json:"card_no"               orm:"card_no"               description:"卡号"`                    // 卡号
	DepositBankProvince string      `json:"deposit_bank_province" orm:"deposit_bank_province" description:"开户行所在省"`                // 开户行所在省
	DepositBankCity     string      `json:"deposit_bank_city"     orm:"deposit_bank_city"     description:"开户行所在市"`                // 开户行所在市
	DepositBank         string      `json:"deposit_bank"          orm:"deposit_bank"          description:"开户行"`                   // 开户行
	RegisterIp          string      `json:"register_ip"           orm:"register_ip"           description:""`                      //
	RegisterTime        *gtime.Time `json:"register_time"         orm:"register_time"         description:""`                      //
	LastLoginIp         string      `json:"last_login_ip"         orm:"last_login_ip"         description:""`                      //
	LastLoginTime       *gtime.Time `json:"last_login_time"       orm:"last_login_time"       description:""`                      //
	CreatedAt           *gtime.Time `json:"created_at"            orm:"created_at"            description:""`                      //
	UpdatedAt           *gtime.Time `json:"updated_at"            orm:"updated_at"            description:""`                      //
	Reason              string      `json:"reason"                orm:"reason"                description:"申请理由"`                  // 申请理由
	Agentcode           string      `json:"agentcode"             orm:"agentcode"             description:""`                      //
}
