// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserGrade is the golang structure for table user_grade.
type UserGrade struct {
	Id                   uint        `json:"id"                   orm:"id"                     description:""`              //
	SiteId               int         `json:"siteId"               orm:"site_id"                description:"站点ID"`          // 站点ID
	Name                 string      `json:"name"                 orm:"name"                   description:"会员等级名称"`        // 会员等级名称
	DepositAmount        float64     `json:"depositAmount"        orm:"deposit_amount"         description:"充值金额"`          // 充值金额
	BetAmount            float64     `json:"betAmount"            orm:"bet_amount"             description:"投注金额"`          // 投注金额
	PointsUpgrade        int         `json:"pointsUpgrade"        orm:"points_upgrade"         description:"升级消耗积分"`        // 升级消耗积分
	BonusUpgrade         float64     `json:"bonusUpgrade"         orm:"bonus_upgrade"          description:"额外返水比例: 体育"`    // 额外返水比例: 体育
	BonusBirthday        float64     `json:"bonusBirthday"        orm:"bonus_birthday"         description:"生日彩金"`          // 生日彩金
	RebatePercentSports  float64     `json:"rebatePercentSports"  orm:"rebate_percent_sports"  description:"额外返水比例: 体育"`    // 额外返水比例: 体育
	RebatePercentLottery float64     `json:"rebatePercentLottery" orm:"rebate_percent_lottery" description:"额外返水比例: 彩票"`    // 额外返水比例: 彩票
	RebatePercentLive    float64     `json:"rebatePercentLive"    orm:"rebate_percent_live"    description:"额外返水比例: 真人视讯"`  // 额外返水比例: 真人视讯
	RebatePercentEgame   float64     `json:"rebatePercentEgame"   orm:"rebate_percent_egame"   description:"额外返水比例: 电子游戏"`  // 额外返水比例: 电子游戏
	RebatePercentPoker   float64     `json:"rebatePercentPoker"   orm:"rebate_percent_poker"   description:"额外返水比例：棋牌"`     // 额外返水比例：棋牌
	FieldsDisable        string      `json:"fieldsDisable"        orm:"fields_disable"         description:"字段开关，用来关闭哪些字段"` // 字段开关，用来关闭哪些字段
	AutoProviding        string      `json:"autoProviding"        orm:"auto_providing"         description:"哪些字段的业务是自动发放的"` // 哪些字段的业务是自动发放的
	Status               int         `json:"status"               orm:"status"                 description:"状态.1=可用；0=禁用"`  // 状态.1=可用；0=禁用
	PaymentLimit         float64     `json:"paymentLimit"         orm:"payment_limit"          description:"充值要求"`          // 充值要求
	BetLimit             float64     `json:"betLimit"             orm:"bet_limit"              description:"投注要求"`          // 投注要求
	MoneyLimit           float64     `json:"moneyLimit"           orm:"money_limit"            description:"升级送"`           // 升级送
	PaymentLimitMonth    float64     `json:"paymentLimitMonth"    orm:"payment_limit_month"    description:"月礼金充值要求"`       // 月礼金充值要求
	BetLimitMonth        float64     `json:"betLimitMonth"        orm:"bet_limit_month"        description:"月礼金投注要求"`       // 月礼金投注要求
	MoneyMonth           float64     `json:"moneyMonth"           orm:"money_month"            description:"月礼金送"`          // 月礼金送
	PaymentLimitWeek     float64     `json:"paymentLimitWeek"     orm:"payment_limit_week"     description:"周礼金充值要求"`       // 周礼金充值要求
	BetLimitWeek         float64     `json:"betLimitWeek"         orm:"bet_limit_week"         description:"周礼金投注要求"`       // 周礼金投注要求
	MoneyWeek            float64     `json:"moneyWeek"            orm:"money_week"             description:"周礼金送"`          // 周礼金送
	DailyWithdrawTimes   int         `json:"dailyWithdrawTimes"   orm:"daily_withdraw_times"   description:"每日提现次数"`        // 每日提现次数
	CreatedAt            *gtime.Time `json:"createdAt"            orm:"created_at"             description:""`              //
	UpdatedAt            *gtime.Time `json:"updatedAt"            orm:"updated_at"             description:""`              //
}
