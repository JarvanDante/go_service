// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RechargeModule is the golang structure for table recharge_module.
type RechargeModule struct {
	Id                     uint        `json:"id"                        orm:"id"                        description:""`                                         //
	SiteId                 int         `json:"site_id"                   orm:"site_id"                   description:"站点ID"`                                     // 站点ID
	ActivityId             int         `json:"activity_id"               orm:"activity_id"               description:"活动id"`                                     // 活动id
	Name                   string      `json:"name"                      orm:"name"                      description:"活动名称"`                                     // 活动名称
	StartTime              *gtime.Time `json:"start_time"                orm:"start_time"                description:"开始时间"`                                     // 开始时间
	EndTime                *gtime.Time `json:"end_time"                  orm:"end_time"                  description:"结束时间"`                                     // 结束时间
	Status                 int         `json:"status"                    orm:"status"                    description:"活动状态。1=开启；0=关闭"`                           // 活动状态。1=开启；0=关闭
	UserGradeIds           string      `json:"user_grade_ids"            orm:"user_grade_ids"            description:"会员等级。以,的形式隔开"`                             // 会员等级。以,的形式隔开
	UserLevelIds           string      `json:"user_level_ids"            orm:"user_level_ids"            description:"会员层级。以,的形式隔开"`                             // 会员层级。以,的形式隔开
	Platform               int         `json:"platform"                  orm:"platform"                  description:"优惠终端。0=所有；1=网站；2=手机"`                      // 优惠终端。0=所有；1=网站；2=手机
	PaymentGateways        string      `json:"payment_gateways"          orm:"payment_gateways"          description:"在线支付网关。以,的形式隔开"`                           // 在线支付网关。以,的形式隔开
	TransferTypes          string      `json:"transfer_types"            orm:"transfer_types"            description:"转账汇款类型。以,的形式隔开"`                           // 转账汇款类型。以,的形式隔开
	UserType               int         `json:"user_type"                 orm:"user_type"                 description:"面向用户类型。0=所有用户，1=新用户；2=老用户"`                // 面向用户类型。0=所有用户，1=新用户；2=老用户
	ConditionRule          int         `json:"condition_rule"            orm:"condition_rule"            description:"充值条件。1=当日单笔充值，2=本周单笔充值；3=本月单笔充值；4=会员首次充值"` // 充值条件。1=当日单笔充值，2=本周单笔充值；3=本月单笔充值；4=会员首次充值
	ConditionCompare       int         `json:"condition_compare"         orm:"condition_compare"         description:"充值条件。1=大于等于，2=等于"`                         // 充值条件。1=大于等于，2=等于
	ConditionValue         float64     `json:"condition_value"           orm:"condition_value"           description:"充值条件。值"`                                   // 充值条件。值
	BonusType              int         `json:"bonus_type"                orm:"bonus_type"                description:"奖励金额类型。1=奖励实际金额；0=奖励金额比例"`                 // 奖励金额类型。1=奖励实际金额；0=奖励金额比例
	BonusAmount            float64     `json:"bonus_amount"              orm:"bonus_amount"              description:"奖励金额"`                                     // 奖励金额
	BonusPercent           float64     `json:"bonus_percent"             orm:"bonus_percent"             description:"奖励金额比例"`                                   // 奖励金额比例
	BonusMax               float64     `json:"bonus_max"                 orm:"bonus_max"                 description:"最高奖励金额.0=不限制"`                             // 最高奖励金额.0=不限制
	ApplyFrequency         int         `json:"apply_frequency"           orm:"apply_frequency"           description:"申请频率。0=不限；1=每日；2=每周；3=每月"`                 // 申请频率。0=不限；1=每日；2=每周；3=每月
	ApplyTimes             int         `json:"apply_times"               orm:"apply_times"               description:"每人申请次数"`                                   // 每人申请次数
	WithdrawNeedWaterTimes int         `json:"withdraw_need_water_times" orm:"withdraw_need_water_times" description:"提现流水要求多少倍流水"`                              // 提现流水要求多少倍流水
	CreatedAt              *gtime.Time `json:"created_at"                orm:"created_at"                description:""`                                         //
	UpdatedAt              *gtime.Time `json:"updated_at"                orm:"updated_at"                description:""`                                         //
}
