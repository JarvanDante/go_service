// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// BalanceLogWithdraw is the golang structure for table balance_log_withdraw.
type BalanceLogWithdraw struct {
	Id          uint64      `json:"id"            orm:"id"            description:""`                    //
	SiteId      int         `json:"site_id"       orm:"site_id"       description:"站点ID"`                // 站点ID
	UserId      int         `json:"user_id"       orm:"user_id"       description:"会员ID"`                // 会员ID
	UserLevelId int         `json:"user_level_id" orm:"user_level_id" description:"会员层级ID"`              // 会员层级ID
	Username    string      `json:"username"      orm:"username"      description:"会员用户名"`               // 会员用户名
	TradeNo     string      `json:"trade_no"      orm:"trade_no"      description:"流水号"`                 // 流水号
	Money       float64     `json:"money"         orm:"money"         description:"充值金额"`                // 充值金额
	Fee         float64     `json:"fee"           orm:"fee"           description:"手续费"`                 // 手续费
	BankName    string      `json:"bank_name"     orm:"bank_name"     description:"提现银行"`                // 提现银行
	CardAccount string      `json:"card_account"  orm:"card_account"  description:"银行户名"`                // 银行户名
	CardNo      string      `json:"card_no"       orm:"card_no"       description:"卡号"`                  // 卡号
	DepositBank string      `json:"deposit_bank"  orm:"deposit_bank"  description:"开户行"`                 // 开户行
	Status      int         `json:"status"        orm:"status"        description:"状态. 1=待审核；2=成功；3=失败"` // 状态. 1=待审核；2=成功；3=失败
	AdminId     int         `json:"admin_id"      orm:"admin_id"      description:"后台用户ID"`              // 后台用户ID
	Remark      string      `json:"remark"        orm:"remark"        description:"备注"`                  // 备注
	CreatedAt   *gtime.Time `json:"created_at"    orm:"created_at"    description:""`                    //
	UpdatedAt   *gtime.Time `json:"updated_at"    orm:"updated_at"    description:""`                    //
}
