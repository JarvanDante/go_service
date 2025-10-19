// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for the table user.
type UserDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  UserColumns        // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// UserColumns defines and stores column names for the table user.
type UserColumns struct {
	Id                string //
	SiteId            string //
	GradeId           string //
	LevelId           string //
	AgentId           string //
	ChannelId         string // 渠道id
	Username          string //
	Password          string //
	PayPassword       string //
	Status            string //
	RegisterIp        string //
	RegisterTime      string //
	RegisterUrl       string // 注册来源
	RegisterDevice    string // 1;电脑2; 手机3;平台
	DeviceCode        string // 设备码
	LastLoginIp       string //
	LastLoginTime     string //
	LastLoginAddress  string // 最后登录地址
	Realname          string //
	Mobile            string //
	Email             string // 邮箱
	Qq                string // QQ号
	Birthday          string //
	Sex               string // 性别。0=未知；1=男；2=女
	IsOnline          string //
	FocusLevel        string // 会员关注级别。1=正常；2=可疑；3=危险
	BalanceStatus     string // 1=0=
	SafeQuestion      string // 密保问题
	SafeAnswer        string // 密保答案
	ShowBeginnerGuide string // 是否显示新手引导。1=显示；0=不显示
	DeleteAt          string // 是否删除。0=未删除；其他为删除时间戳
	Remark            string // 备注
	SiteCode          string // 应用码
	ChannelCode       string // 渠道码
	CreatedAt         string //
	UpdatedAt         string //
	PayTimes          string // 充值次数
	HasPass           string // 是否设置密码
	Currency          string // 用户币种
	WithdrawTimes     string // 提现次数
}

// userColumns holds the columns for the table user.
var userColumns = UserColumns{
	Id:                "id",
	SiteId:            "site_id",
	GradeId:           "grade_id",
	LevelId:           "level_id",
	AgentId:           "agent_id",
	ChannelId:         "channel_id",
	Username:          "username",
	Password:          "password",
	PayPassword:       "pay_password",
	Status:            "status",
	RegisterIp:        "register_ip",
	RegisterTime:      "register_time",
	RegisterUrl:       "register_url",
	RegisterDevice:    "register_device",
	DeviceCode:        "device_code",
	LastLoginIp:       "last_login_ip",
	LastLoginTime:     "last_login_time",
	LastLoginAddress:  "last_login_address",
	Realname:          "realname",
	Mobile:            "mobile",
	Email:             "email",
	Qq:                "qq",
	Birthday:          "birthday",
	Sex:               "sex",
	IsOnline:          "is_online",
	FocusLevel:        "focus_level",
	BalanceStatus:     "balance_status",
	SafeQuestion:      "safe_question",
	SafeAnswer:        "safe_answer",
	ShowBeginnerGuide: "show_beginner_guide",
	DeleteAt:          "delete_at",
	Remark:            "remark",
	SiteCode:          "site_code",
	ChannelCode:       "channel_code",
	CreatedAt:         "created_at",
	UpdatedAt:         "updated_at",
	PayTimes:          "pay_times",
	HasPass:           "has_pass",
	Currency:          "currency",
	WithdrawTimes:     "withdraw_times",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao(handlers ...gdb.ModelHandler) *UserDao {
	return &UserDao{
		group:    "by_site",
		table:    "user",
		columns:  userColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
