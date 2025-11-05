// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// AgentSharingLog is the golang structure for table agent_sharing_log.
type AgentSharingLog struct {
	Id                    uint        `json:"id"                     orm:"id"                     description:""`                             //
	SiteId                int         `json:"site_id"                orm:"site_id"                description:"站点ID"`                         // 站点ID
	StartDate             *gtime.Time `json:"start_date"             orm:"start_date"             description:"开始日期"`                         // 开始日期
	EndDate               *gtime.Time `json:"end_date"               orm:"end_date"               description:"结束日期"`                         // 结束日期
	AgentId               int         `json:"agent_id"               orm:"agent_id"               description:"代理ID"`                         // 代理ID
	AgentUsername         string      `json:"agent_username"         orm:"agent_username"         description:"代理账号"`                         // 代理账号
	UserCount             int         `json:"user_count"             orm:"user_count"             description:"有效会员数"`                        // 有效会员数
	RechargeAmount        float64     `json:"recharge_amount"        orm:"recharge_amount"        description:"充值总计"`                         // 充值总计
	WithdrawAmount        float64     `json:"withdraw_amount"        orm:"withdraw_amount"        description:"提现总计"`                         // 提现总计
	BonusAmount           float64     `json:"bonus_amount"           orm:"bonus_amount"           description:"红利总计"`                         // 红利总计
	RebateAmount          float64     `json:"rebate_amount"          orm:"rebate_amount"          description:"返水总计"`                         // 返水总计
	ValidBetAmount        float64     `json:"valid_bet_amount"       orm:"valid_bet_amount"       description:"有效投注总计"`                       // 有效投注总计
	WinOrLose             float64     `json:"win_or_lose"            orm:"win_or_lose"            description:"输赢总计"`                         // 输赢总计
	FeeAmount             float64     `json:"fee_amount"             orm:"fee_amount"             description:"手续费总计"`                        // 手续费总计
	AdministrationExpense float64     `json:"administration_expense" orm:"administration_expense" description:"行政费用"`                         // 行政费用
	CalculateCommission   float64     `json:"calculate_commission"   orm:"calculate_commission"   description:"程序计算出的代理佣金"`                   // 程序计算出的代理佣金
	Commission            float64     `json:"commission"             orm:"commission"             description:"发放佣金"`                         // 发放佣金
	Status                int         `json:"status"                 orm:"status"                 description:"状态。1=未结算；2=结算成功；3=结算失败；4=未盈利"` // 状态。1=未结算；2=结算成功；3=结算失败；4=未盈利
	AdminId               int         `json:"admin_id"               orm:"admin_id"               description:"后台员工ID"`                       // 后台员工ID
	CreatedAt             *gtime.Time `json:"created_at"             orm:"created_at"             description:""`                             //
	UpdatedAt             *gtime.Time `json:"updated_at"             orm:"updated_at"             description:""`                             //
	CommissionData        string      `json:"commission_data"        orm:"commission_data"        description:"返佣金额计算公式数据"`                   // 返佣金额计算公式数据
}
