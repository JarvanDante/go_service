// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// WaterConfigDao is the data access object for the table water_config.
type WaterConfigDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  WaterConfigColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// WaterConfigColumns defines and stores column names for the table water_config.
type WaterConfigColumns struct {
	Id         string // ID
	Name       string // 配置项名称
	Style      string // 主类型:充值/加款/活动/返水
	Type       string // 子类型:签到/升级/注册送/首充送/大转盘/周礼金/月礼金
	WaterTimes string // 倍水数
	CreatedAt  string // 创建时间
	UpdatedAt  string // 更新时间
}

// waterConfigColumns holds the columns for the table water_config.
var waterConfigColumns = WaterConfigColumns{
	Id:         "id",
	Name:       "name",
	Style:      "style",
	Type:       "type",
	WaterTimes: "water_times",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
}

// NewWaterConfigDao creates and returns a new DAO object for table data access.
func NewWaterConfigDao(handlers ...gdb.ModelHandler) *WaterConfigDao {
	return &WaterConfigDao{
		group:    "by_site",
		table:    "water_config",
		columns:  waterConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *WaterConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *WaterConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *WaterConfigDao) Columns() WaterConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *WaterConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *WaterConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *WaterConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
