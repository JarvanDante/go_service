// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PaymentGatewayDao is the data access object for the table payment_gateway.
type PaymentGatewayDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  PaymentGatewayColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// PaymentGatewayColumns defines and stores column names for the table payment_gateway.
type PaymentGatewayColumns struct {
	Id            string //
	PaymentId     string //
	Gateway       string //
	Status        string //
	MobileSupport string //
	UrlOrder      string //
	UrlQuery      string //
	IsSelectBank  string // 是否需要选择银行
	CreatedAt     string //
	UpdatedAt     string //
}

// paymentGatewayColumns holds the columns for the table payment_gateway.
var paymentGatewayColumns = PaymentGatewayColumns{
	Id:            "id",
	PaymentId:     "payment_id",
	Gateway:       "gateway",
	Status:        "status",
	MobileSupport: "mobile_support",
	UrlOrder:      "url_order",
	UrlQuery:      "url_query",
	IsSelectBank:  "is_select_bank",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewPaymentGatewayDao creates and returns a new DAO object for table data access.
func NewPaymentGatewayDao(handlers ...gdb.ModelHandler) *PaymentGatewayDao {
	return &PaymentGatewayDao{
		group:    "default",
		table:    "payment_gateway",
		columns:  paymentGatewayColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *PaymentGatewayDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *PaymentGatewayDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *PaymentGatewayDao) Columns() PaymentGatewayColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *PaymentGatewayDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *PaymentGatewayDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *PaymentGatewayDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
