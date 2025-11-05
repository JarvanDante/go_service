// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogManual is the golang structure for table balance_log_manual.
type BalanceLogManual struct {
	Id         uint64      `json:"id"          orm:"id"          description:""`                                                                                                                             //
	SiteId     int         `json:"site_id"     orm:"site_id"     description:"站点ID"`                                                                                                                         // 站点ID
	TradeType  int         `json:"trade_type"  orm:"trade_type"  description:"交易类型。1=加款-充值；2=加款-红利；3=加款-返水；4=补款（不入充值记录），5=扣款 - 提现，6=扣款（不入提现记录）7=第三方加款（不入账变）,8=第三方扣款（不入账变）;9=第三方账户 -> 中心账户;10=中心账户 -> 第三方账户"` // 交易类型。1=加款-充值；2=加款-红利；3=加款-返水；4=补款（不入充值记录），5=扣款 - 提现，6=扣款（不入提现记录）7=第三方加款（不入账变）,8=第三方扣款（不入账变）;9=第三方账户 -> 中心账户;10=中心账户 -> 第三方账户
	UserId     int         `json:"user_id"     orm:"user_id"     description:"用户ID"`                                                                                                                         // 用户ID
	Username   string      `json:"username"    orm:"username"    description:"会员用户名"`                                                                                                                        // 会员用户名
	GameId     int         `json:"game_id"     orm:"game_id"     description:"游戏ID"`                                                                                                                         // 游戏ID
	TradeNo    string      `json:"trade_no"    orm:"trade_no"    description:"流水号"`                                                                                                                          // 流水号
	Money      float64     `json:"money"       orm:"money"       description:"金额"`                                                                                                                           // 金额
	Fee        float64     `json:"fee"         orm:"fee"         description:"手续费"`                                                                                                                          // 手续费
	WaterTimes int         `json:"water_times" orm:"water_times" description:"流水倍数"`                                                                                                                         // 流水倍数
	Status     int         `json:"status"      orm:"status"      description:"状态. 1=待审核；2=成功；3=失败"`                                                                                                          // 状态. 1=待审核；2=成功；3=失败
	AdminId    int         `json:"admin_id"    orm:"admin_id"    description:"后台管理员ID"`                                                                                                                      // 后台管理员ID
	Remark     string      `json:"remark"      orm:"remark"      description:"备注"`                                                                                                                           // 备注
	CreatedAt  *gtime.Time `json:"created_at"  orm:"created_at"  description:""`                                                                                                                             //
	UpdatedAt  *gtime.Time `json:"updated_at"  orm:"updated_at"  description:""`                                                                                                                             //
}
