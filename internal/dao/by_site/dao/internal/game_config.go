// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// GameConfigDao is the data access object for the table game_config.
type GameConfigDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  GameConfigColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// GameConfigColumns defines and stores column names for the table game_config.
type GameConfigColumns struct {
	Id        string // ID
	SiteId    string // 应用id
	ChannelId string // 渠道id
	GameId    string // 游戏id
	Vendor    string // 游戏厂商
	Status    string // 状态,1开0关
	Available string // 可用状态1可用2维护
	CreatedAt string // 创建时间
	UpdatedAt string // 更新时间
}

// gameConfigColumns holds the columns for the table game_config.
var gameConfigColumns = GameConfigColumns{
	Id:        "id",
	SiteId:    "site_id",
	ChannelId: "channel_id",
	GameId:    "game_id",
	Vendor:    "vendor",
	Status:    "status",
	Available: "available",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// NewGameConfigDao creates and returns a new DAO object for table data access.
func NewGameConfigDao(handlers ...gdb.ModelHandler) *GameConfigDao {
	return &GameConfigDao{
		group:    "by_site",
		table:    "game_config",
		columns:  gameConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *GameConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *GameConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *GameConfigDao) Columns() GameConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *GameConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *GameConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *GameConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
