// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// BalanceLogPaymentDao is the data access object for the table balance_log_payment.
type BalanceLogPaymentDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  BalanceLogPaymentColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// BalanceLogPaymentColumns defines and stores column names for the table balance_log_payment.
type BalanceLogPaymentColumns struct {
	Id                 string //
	SiteId             string // 站点ID
	ChannelId          string // 渠道id
	UserId             string // 用户ID
	Username           string // 会员用户名
	ActivityRechargeId string // 充值活动ID
	Gateway            string // 网关类型
	PaymentId          string // 支付ID
	PaymentAccountId   string // 支付接口账号ID
	BankValue          string // 银行代码
	TradeNo            string // 流水号
	Money              string // 充值金额
	Fee                string // 手续费
	Status             string // 状态。1=未付；2=成功；3=失败
	AdminId            string // 后台管理员ID。默认为0
	Remark             string // 备注
	Channel            string // 支付渠道
	CardNo             string // 支付卡号
	ImageUrl           string // 渠道引导图片
	SysTradeNo         string // 支付系统订单流水号
	CreatedAt          string //
	UpdatedAt          string //
	Currency           string // 币种
}

// balanceLogPaymentColumns holds the columns for the table balance_log_payment.
var balanceLogPaymentColumns = BalanceLogPaymentColumns{
	Id:                 "id",
	SiteId:             "site_id",
	ChannelId:          "channel_id",
	UserId:             "user_id",
	Username:           "username",
	ActivityRechargeId: "activity_recharge_id",
	Gateway:            "gateway",
	PaymentId:          "payment_id",
	PaymentAccountId:   "payment_account_id",
	BankValue:          "bank_value",
	TradeNo:            "trade_no",
	Money:              "money",
	Fee:                "fee",
	Status:             "status",
	AdminId:            "admin_id",
	Remark:             "remark",
	Channel:            "channel",
	CardNo:             "card_no",
	ImageUrl:           "image_url",
	SysTradeNo:         "sys_trade_no",
	CreatedAt:          "created_at",
	UpdatedAt:          "updated_at",
	Currency:           "currency",
}

// NewBalanceLogPaymentDao creates and returns a new DAO object for table data access.
func NewBalanceLogPaymentDao(handlers ...gdb.ModelHandler) *BalanceLogPaymentDao {
	return &BalanceLogPaymentDao{
		group:    "by_balance",
		table:    "balance_log_payment",
		columns:  balanceLogPaymentColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *BalanceLogPaymentDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *BalanceLogPaymentDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *BalanceLogPaymentDao) Columns() BalanceLogPaymentColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *BalanceLogPaymentDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *BalanceLogPaymentDao) Ctx(ctx context.Context) *gdb.Model {
	model := dao.DB().Model(dao.table)
	for _, handler := range dao.handlers {
		model = handler(model)
	}
	return model.Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rolls back the transaction and returns the error if function f returns a non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note: Do not commit or roll back the transaction in function f,
// as it is automatically handled by this function.
func (dao *BalanceLogPaymentDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
