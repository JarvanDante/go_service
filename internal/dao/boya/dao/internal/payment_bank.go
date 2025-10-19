// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PaymentBankDao is the data access object for the table payment_bank.
type PaymentBankDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  PaymentBankColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// PaymentBankColumns defines and stores column names for the table payment_bank.
type PaymentBankColumns struct {
	Id        string //
	PaymentId string // 支付ID
	BankId    string // 银行ID
	BankName  string // 银行名称
	BankCode  string // 银行标识
	BankValue string // 支付接口银行code
	Status    string // 状态。1=可用；0=禁用
	Sort      string // 排序
	CreatedAt string //
	UpdatedAt string //
}

// paymentBankColumns holds the columns for the table payment_bank.
var paymentBankColumns = PaymentBankColumns{
	Id:        "id",
	PaymentId: "payment_id",
	BankId:    "bank_id",
	BankName:  "bank_name",
	BankCode:  "bank_code",
	BankValue: "bank_value",
	Status:    "status",
	Sort:      "sort",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewPaymentBankDao creates and returns a new DAO object for table data access.
func NewPaymentBankDao(handlers ...gdb.ModelHandler) *PaymentBankDao {
	return &PaymentBankDao{
		group:    "default",
		table:    "payment_bank",
		columns:  paymentBankColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PaymentBankDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PaymentBankDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PaymentBankDao) Columns() PaymentBankColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PaymentBankDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PaymentBankDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PaymentBankDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
