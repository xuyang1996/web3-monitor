package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ TransferTxnModel = (*customTransferTxnModel)(nil)

type (
	// TransferTxnModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTransferTxnModel.
	TransferTxnModel interface {
		transferTxnModel
	}

	customTransferTxnModel struct {
		*defaultTransferTxnModel
	}
)

// NewTransferTxnModel returns a model for the database table.
func NewTransferTxnModel(conn sqlx.SqlConn) TransferTxnModel {
	return &customTransferTxnModel{
		defaultTransferTxnModel: newTransferTxnModel(conn),
	}
}
