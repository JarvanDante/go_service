// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AgSportsDao is the data access object for the table ag_sports.
type AgSportsDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  AgSportsColumns    // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// AgSportsColumns defines and stores column names for the table ag_sports.
type AgSportsColumns struct {
	Id                    string // ID
	GameNo                string // 用于标识此下注请求交易的唯一ID。
	InningNo              string // 用于将所有下注和赢取分组在单个回合中的游戏回合ID
	ExternalTransactionId string // 由游戏供应商提供的外部交易ID
	Username              string // 用户名
	CurrencyCode          string // 币种
	GameCode              string // OneAPI集成系统中选定游戏的游戏代码
	VendorCode            string // 提供商code
	GameCategoryCode      string // 游戏分类code
	BetAmount             string // 投注金额
	WinAmount             string // 赢取的金额，当winAmount金额>0时需要钱包贷记操作。
	Win                   string // 绝对赢取或亏损的金额
	ValidAmount           string // 有效投注额的金额
	JackpotAmount         string // 奖池金额，当jackpotAmount金额>0时需要钱包贷记操作
	BetStatus             string //
	BetTime               string //
	VendorSettleTime      string //
	IsFreeSpin            string // 用于指示下注为免费旋转下注的状态
	VendorBetId           string //
	CreatedAt             string //
	UpdatedAt             string //
}

// agSportsColumns holds the columns for the table ag_sports.
var agSportsColumns = AgSportsColumns{
	Id:                    "id",
	GameNo:                "game_no",
	InningNo:              "inning_no",
	ExternalTransactionId: "externalTransactionId",
	Username:              "username",
	CurrencyCode:          "currencyCode",
	GameCode:              "gameCode",
	VendorCode:            "vendorCode",
	GameCategoryCode:      "gameCategoryCode",
	BetAmount:             "bet_amount",
	WinAmount:             "winAmount",
	Win:                   "win",
	ValidAmount:           "valid_amount",
	JackpotAmount:         "jackpotAmount",
	BetStatus:             "bet_status",
	BetTime:               "bet_time",
	VendorSettleTime:      "vendorSettleTime",
	IsFreeSpin:            "isFreeSpin",
	VendorBetId:           "vendorBetId",
	CreatedAt:             "created_at",
	UpdatedAt:             "updated_at",
}

// NewAgSportsDao creates and returns a new DAO object for table data access.
func NewAgSportsDao(handlers ...gdb.ModelHandler) *AgSportsDao {
	return &AgSportsDao{
		group:    "casion_by",
		table:    "ag_sports",
		columns:  agSportsColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *AgSportsDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *AgSportsDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *AgSportsDao) Columns() AgSportsColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *AgSportsDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *AgSportsDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *AgSportsDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
