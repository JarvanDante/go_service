// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// RiskForewarnConfig is the golang structure for table risk_forewarn_config.
type RiskForewarnConfig struct {
	Id                    uint64      `json:"id"                       orm:"id"                       description:""`                                 //
	SiteId                int         `json:"site_id"                  orm:"site_id"                  description:"站点ID"`                             // 站点ID
	AccountToGame         float64     `json:"account_to_game"          orm:"account_to_game"          description:"账户转账到游戏单笔金额预警, 0=不提示 其他=提示"`       // 账户转账到游戏单笔金额预警, 0=不提示 其他=提示
	GameToAccount         float64     `json:"game_to_account"          orm:"game_to_account"          description:"游戏转账到账户单笔金额预警, 0=不提示 其他=提示"`       // 游戏转账到账户单笔金额预警, 0=不提示 其他=提示
	WinMoney              float64     `json:"win_money"                orm:"win_money"                description:"会员单笔游戏赢得金额预警,0=不提示 其他=提示"`         // 会员单笔游戏赢得金额预警,0=不提示 其他=提示
	BetAmount             float64     `json:"bet_amount"               orm:"bet_amount"               description:"会员单笔游戏投注金额预警,0=不提示 其他=提示"`         // 会员单笔游戏投注金额预警,0=不提示 其他=提示
	AlterBankCard         int         `json:"alter_bank_card"          orm:"alter_bank_card"          description:"会员银行卡号被修改预警，0=不提示；1=提示"`           // 会员银行卡号被修改预警，0=不提示；1=提示
	LoginAreaDifference   int         `json:"login_area_difference"    orm:"login_area_difference"    description:"本次和上次登录异样提示，0=不提示；1=提示"`           // 本次和上次登录异样提示，0=不提示；1=提示
	SameipRegisterTime    int         `json:"sameip_register_time"     orm:"sameip_register_time"     description:"同一IP单位时间内注册限制，0=不限时间 其他=时间，单位：分钟"` // 同一IP单位时间内注册限制，0=不限时间 其他=时间，单位：分钟
	SameipRegisterCount   int         `json:"sameip_register_count"    orm:"sameip_register_count"    description:"同一IP单位时间内注册次数限制，0=不限 其他=提示"`       // 同一IP单位时间内注册次数限制，0=不限 其他=提示
	CreatedAt             *gtime.Time `json:"created_at"               orm:"created_at"               description:""`                                 //
	UpdatedAt             *gtime.Time `json:"updated_at"               orm:"updated_at"               description:""`                                 //
	IsAccountToGame       int         `json:"is_account_to_game"       orm:"is_account_to_game"       description:"是否开启会员账户-游戏预警"`                    // 是否开启会员账户-游戏预警
	IsGameToAccount       int         `json:"is_game_to_account"       orm:"is_game_to_account"       description:"是否开启会员游戏-账户预警"`                    // 是否开启会员游戏-账户预警
	IsWinMoney            int         `json:"is_win_money"             orm:"is_win_money"             description:"是否开启游戏输赢预警"`                       // 是否开启游戏输赢预警
	IsBetAmount           int         `json:"is_bet_amount"            orm:"is_bet_amount"            description:"是否开启游戏投注预警"`                       // 是否开启游戏投注预警
	IsAlterBankCard       int         `json:"is_alter_bank_card"       orm:"is_alter_bank_card"       description:"是否开启修改银行卡预警"`                      // 是否开启修改银行卡预警
	IsLoginAreaDifference int         `json:"is_login_area_difference" orm:"is_login_area_difference" description:"是否开启登录地区不同预警"`                     // 是否开启登录地区不同预警
	IsSameipRegisterCount int         `json:"is_sameip_register_count" orm:"is_sameip_register_count" description:"是否开启相同IP注册次数预警"`                   // 是否开启相同IP注册次数预警
}
