// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserMessageDao is the data access object for the table user_message.
type UserMessageDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserMessageColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserMessageColumns defines and stores column names for the table user_message.
type UserMessageColumns struct {
	Id        string //
	SiteId    string // 站点ID
	UserId    string // 用户ID
	MessageId string // 消息ID
	Status    string // 状态。1=已读；0=未读
	DeleteAt  string // 是否删除。0=未删除；其他为删除时间戳
	CreatedAt string //
	UpdatedAt string //
}

// userMessageColumns holds the columns for the table user_message.
var userMessageColumns = UserMessageColumns{
	Id:        "id",
	SiteId:    "site_id",
	UserId:    "user_id",
	MessageId: "message_id",
	Status:    "status",
	DeleteAt:  "delete_at",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewUserMessageDao creates and returns a new DAO object for table data access.
func NewUserMessageDao(handlers ...gdb.ModelHandler) *UserMessageDao {
	return &UserMessageDao{
		group:    "by_site",
		table:    "user_message",
		columns:  userMessageColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserMessageDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserMessageDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserMessageDao) Columns() UserMessageColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserMessageDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserMessageDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserMessageDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
