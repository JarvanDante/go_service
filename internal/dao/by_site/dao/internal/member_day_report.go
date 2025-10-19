// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// MemberDayReportDao is the data access object for the table member_day_report.
type MemberDayReportDao struct {
	table    string                 // table is the underlying table name of the DAO.
	group    string                 // group is the database configuration group name of the current DAO.
	columns  MemberDayReportColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler     // handlers for customized model modification.
}

// MemberDayReportColumns defines and stores column names for the table member_day_report.
type MemberDayReportColumns struct {
	Id            string // 主键ID
	SiteId        string // 商户ID
	UserId        string // 会员ID
	Date          string // 统计日期
	DepositAmount string // 当日充值总额
	BetAmount     string // 当日投注总额
	CreatedAt     string // 记录创建时间
	UpdatedAt     string // 记录更新时间
}

// memberDayReportColumns holds the columns for the table member_day_report.
var memberDayReportColumns = MemberDayReportColumns{
	Id:            "id",
	SiteId:        "site_id",
	UserId:        "user_id",
	Date:          "date",
	DepositAmount: "deposit_amount",
	BetAmount:     "bet_amount",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
}

// NewMemberDayReportDao creates and returns a new DAO object for table data access.
func NewMemberDayReportDao(handlers ...gdb.ModelHandler) *MemberDayReportDao {
	return &MemberDayReportDao{
		group:    "by_site",
		table:    "member_day_report",
		columns:  memberDayReportColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *MemberDayReportDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *MemberDayReportDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *MemberDayReportDao) Columns() MemberDayReportColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *MemberDayReportDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *MemberDayReportDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *MemberDayReportDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
