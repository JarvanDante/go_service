// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// ApiAgentConfigDao is the data access object for the table api_agent_config.
type ApiAgentConfigDao struct {
	table    string                // table is the underlying table name of the DAO.
	group    string                // group is the database configuration group name of the current DAO.
	columns  ApiAgentConfigColumns // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler    // handlers for customized model modification.
}

// ApiAgentConfigColumns defines and stores column names for the table api_agent_config.
type ApiAgentConfigColumns struct {
	Id            string //
	AgentName     string //
	AppId         string //
	AppKey        string //
	PayloadSecret string //
	ApiUrl        string //
	CreatedAt     string //
	UpdatedAt     string //
	GameId        string // 厅ID
	Currency      string // 币种
}

// apiAgentConfigColumns holds the columns for the table api_agent_config.
var apiAgentConfigColumns = ApiAgentConfigColumns{
	Id:            "id",
	AgentName:     "agent_name",
	AppId:         "app_id",
	AppKey:        "app_key",
	PayloadSecret: "payload_secret",
	ApiUrl:        "api_url",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	GameId:        "game_id",
	Currency:      "currency",
}

// NewApiAgentConfigDao creates and returns a new DAO object for table data access.
func NewApiAgentConfigDao(handlers ...gdb.ModelHandler) *ApiAgentConfigDao {
	return &ApiAgentConfigDao{
		group:    "by_site",
		table:    "api_agent_config",
		columns:  apiAgentConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *ApiAgentConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *ApiAgentConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *ApiAgentConfigDao) Columns() ApiAgentConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *ApiAgentConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *ApiAgentConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *ApiAgentConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
