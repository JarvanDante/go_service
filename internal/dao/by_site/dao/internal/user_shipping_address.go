// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserShippingAddressDao is the data access object for the table user_shipping_address.
type UserShippingAddressDao struct {
	table    string                     // table is the underlying table name of the DAO.
	group    string                     // group is the database configuration group name of the current DAO.
	columns  UserShippingAddressColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler         // handlers for customized model modification.
}

// UserShippingAddressColumns defines and stores column names for the table user_shipping_address.
type UserShippingAddressColumns struct {
	Id        string //
	SiteId    string // 站点ID
	UserId    string // 用户ID
	Province  string // 省份
	City      string // 城市
	Town      string // 区镇
	Address   string // 详细地址
	Addressee string // 收件人
	Mobile    string // 手机号
	Phone     string // 固定电话
	IsDefault string // 是否是默认收货地址。1=是；0=否
	CreatedAt string //
	UpdatedAt string //
}

// userShippingAddressColumns holds the columns for the table user_shipping_address.
var userShippingAddressColumns = UserShippingAddressColumns{
	Id:        "id",
	SiteId:    "site_id",
	UserId:    "user_id",
	Province:  "province",
	City:      "city",
	Town:      "town",
	Address:   "address",
	Addressee: "addressee",
	Mobile:    "mobile",
	Phone:     "phone",
	IsDefault: "is_default",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewUserShippingAddressDao creates and returns a new DAO object for table data access.
func NewUserShippingAddressDao(handlers ...gdb.ModelHandler) *UserShippingAddressDao {
	return &UserShippingAddressDao{
		group:    "by_site",
		table:    "user_shipping_address",
		columns:  userShippingAddressColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserShippingAddressDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserShippingAddressDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserShippingAddressDao) Columns() UserShippingAddressColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserShippingAddressDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserShippingAddressDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserShippingAddressDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
