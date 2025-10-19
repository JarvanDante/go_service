// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SignConfigDao is the data access object for the table sign_config.
type SignConfigDao struct {
	table    string             // table is the underlying table name of the DAO.
	group    string             // group is the database configuration group name of the current DAO.
	columns  SignConfigColumns  // columns contains all the column names of Table for convenient usage.
	handlers []gdb.ModelHandler // handlers for customized model modification.
}

// SignConfigColumns defines and stores column names for the table sign_config.
type SignConfigColumns struct {
	Id            string // ID
	SiteId        string // 商户id
	Type          string //
	First         string //
	Second        string //
	Third         string //
	Fourth        string //
	Fifth         string //
	Sixth         string //
	Seventh       string //
	Eighth        string //
	Ninth         string //
	Tenth         string //
	Eleventh      string //
	Twelfth       string //
	Thirteenth    string //
	Fourteenth    string //
	Fifteenth     string //
	Sixteenth     string //
	Seventeenth   string //
	Eighteenth    string //
	Nineteenth    string //
	Twenty        string //
	TwentyFirst   string //
	TwentySecond  string //
	TwentyThird   string //
	TwentyFourth  string //
	TwentyFifth   string //
	TwentySixth   string //
	TwentySeventh string //
	TwentyEighth  string //
	TwentyNinth   string //
	Thirtieth     string //
	ThirtyFirst   string //
}

// signConfigColumns holds the columns for the table sign_config.
var signConfigColumns = SignConfigColumns{
	Id:            "id",
	SiteId:        "site_id",
	Type:          "type",
	First:         "first",
	Second:        "second",
	Third:         "third",
	Fourth:        "fourth",
	Fifth:         "fifth",
	Sixth:         "sixth",
	Seventh:       "seventh",
	Eighth:        "eighth",
	Ninth:         "ninth",
	Tenth:         "tenth",
	Eleventh:      "eleventh",
	Twelfth:       "twelfth",
	Thirteenth:    "thirteenth",
	Fourteenth:    "fourteenth",
	Fifteenth:     "fifteenth",
	Sixteenth:     "sixteenth",
	Seventeenth:   "seventeenth",
	Eighteenth:    "eighteenth",
	Nineteenth:    "nineteenth",
	Twenty:        "twenty",
	TwentyFirst:   "twenty_first",
	TwentySecond:  "twenty_second",
	TwentyThird:   "twenty_third",
	TwentyFourth:  "twenty_fourth",
	TwentyFifth:   "twenty_fifth",
	TwentySixth:   "twenty_sixth",
	TwentySeventh: "twenty_seventh",
	TwentyEighth:  "twenty_eighth",
	TwentyNinth:   "twenty_ninth",
	Thirtieth:     "thirtieth",
	ThirtyFirst:   "thirty_first",
}

// NewSignConfigDao creates and returns a new DAO object for table data access.
func NewSignConfigDao(handlers ...gdb.ModelHandler) *SignConfigDao {
	return &SignConfigDao{
		group:    "by_site",
		table:    "sign_config",
		columns:  signConfigColumns,
		handlers: handlers,
	}
}

// DB retrieves and returns the underlying raw database management object of the current DAO.
func (dao *SignConfigDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of the current DAO.
func (dao *SignConfigDao) Table() string {
	return dao.table
}

// Columns returns all column names of the current DAO.
func (dao *SignConfigDao) Columns() SignConfigColumns {
	return dao.columns
}

// Group returns the database configuration group name of the current DAO.
func (dao *SignConfigDao) Group() string {
	return dao.group
}

// Ctx creates and returns a Model for the current DAO. It automatically sets the context for the current operation.
func (dao *SignConfigDao) Ctx(ctx context.Context) *gdb.Model {
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
func (dao *SignConfigDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
