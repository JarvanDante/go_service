// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MemberBonusRecordDao is the data access object for the table member_bonus_record.
type MemberBonusRecordDao struct {
	table    string                   // table is the underlying table name of the DAO.
	group    string                   // group is the database configuration group name of the current DAO.
	columns  MemberBonusRecordColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler       // handlers for customized model modification.
}

// MemberBonusRecordColumns defines and stores column names for the table member_bonus_record.
type MemberBonusRecordColumns struct {
	Id        string // 主键ID
	SiteId    string // 商户ID
	UserId    string // 会员ID
	BonusType string // 礼金类型: 1=周礼金 2=月礼金 3=升级礼金
	Season    string // 周期标识: 周/2025W36, 月/2025-08, 升级/VIP等级
	Amount    string // 领取金额
	ExpiredAt string // 过期时间，升级礼金永不过期可为NULL
	CreatedAt string // 领取时间
	UpdatedAt string // 更新时间
}

// memberBonusRecordColumns holds the columns for the table member_bonus_record.
var memberBonusRecordColumns = MemberBonusRecordColumns{
	Id:        "id",
	SiteId:    "site_id",
	UserId:    "user_id",
	BonusType: "bonus_type",
	Season:    "season",
	Amount:    "amount",
	ExpiredAt: "expired_at",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewMemberBonusRecordDao creates and returns a new DAO object for table data access.
func NewMemberBonusRecordDao(handlers ...gdb.ModelHandler) *MemberBonusRecordDao {
	return &MemberBonusRecordDao{
		group:    "by_site",
		table:    "member_bonus_record",
		columns:  memberBonusRecordColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MemberBonusRecordDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MemberBonusRecordDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MemberBonusRecordDao) Columns() MemberBonusRecordColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MemberBonusRecordDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MemberBonusRecordDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MemberBonusRecordDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
