// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLog is the golang structure for table balance_log.
type BalanceLog struct {
	Id        uint        `json:"id"         orm:"id"         description:""`                         //
	SiteId    int         `json:"site_id"    orm:"site_id"    description:"站点ID"`                     // 站点ID
	UserId    int         `json:"user_id"    orm:"user_id"    description:"会员ID"`                     // 会员ID
	TradeType int         `json:"trade_type" orm:"trade_type" description:"交易类型"`                     // 交易类型
	Gateway   int         `json:"gateway"    orm:"gateway"    description:"支付相关：支付类型"`                // 支付相关：支付类型
	BankType  int         `json:"bank_type"  orm:"bank_type"  description:"转账相关：银行类型"`                // 转账相关：银行类型
	GameId    int         `json:"game_id"    orm:"game_id"    description:"游戏相关：游戏ID"`                // 游戏相关：游戏ID
	TradeNo   string      `json:"trade_no"   orm:"trade_no"   description:"交易流水号"`                    // 交易流水号
	Money     float64     `json:"money"      orm:"money"      description:"金额"`                       // 金额
	Fee       float64     `json:"fee"        orm:"fee"        description:"手续费"`                      // 手续费
	Status    int         `json:"status"     orm:"status"     description:"状态。0=待处理，1=处理中；2=成功；3=失败"` // 状态。0=待处理，1=处理中；2=成功；3=失败
	AdminId   int         `json:"admin_id"   orm:"admin_id"   description:"后台员工ID"`                   // 后台员工ID
	Remark    string      `json:"remark"     orm:"remark"     description:"备注"`                       // 备注
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`                         //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`                         //
}
