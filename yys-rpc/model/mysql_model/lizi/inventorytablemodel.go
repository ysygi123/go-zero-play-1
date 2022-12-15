package lizi

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ InventoryTableModel = (*customInventoryTableModel)(nil)

type (
	// InventoryTableModel is an interface to be customized, add more methods here,
	// and implement the added methods in customInventoryTableModel.
	InventoryTableModel interface {
		inventoryTableModel
	}

	customInventoryTableModel struct {
		*defaultInventoryTableModel
	}
)

// NewInventoryTableModel returns a model for the database table.
func NewInventoryTableModel(conn sqlx.SqlConn, c cache.CacheConf) InventoryTableModel {
	return &customInventoryTableModel{
		defaultInventoryTableModel: newInventoryTableModel(conn, c),
	}
}
