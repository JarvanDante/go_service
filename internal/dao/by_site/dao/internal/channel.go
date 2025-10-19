// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ChannelDao is the data access object for the table channel.
type ChannelDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  ChannelColumns     // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// ChannelColumns defines and stores column names for the table channel.
type ChannelColumns struct {
	Id        string // ID
	SiteId    string // 商户id
	Code      string // 渠道编码
	Name      string // 渠道名称
	Type      string // 渠道类型
	Status    string // 状态0关1开
	Remark    string // 备注
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
	DeletedAt string //
}

// channelColumns holds the columns for the table channel.
var channelColumns = ChannelColumns{
	Id:        "id",
	SiteId:    "site_id",
	Code:      "code",
	Name:      "name",
	Type:      "type",
	Status:    "status",
	Remark:    "remark",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
	DeletedAt: "deleted_at",
}

// NewChannelDao creates and returns a new DAO object for table data access.
func NewChannelDao(handlers ...gdb.ModelHandler) *ChannelDao {
	return &ChannelDao{
		group:    "by_site",
		table:    "channel",
		columns:  channelColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ChannelDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ChannelDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ChannelDao) Columns() ChannelColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ChannelDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ChannelDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ChannelDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
