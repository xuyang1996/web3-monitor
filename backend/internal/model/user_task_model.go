package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserTaskModel = (*customUserTaskModel)(nil)

type (
	// UserTaskModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserTaskModel.
	UserTaskModel interface {
		userTaskModel
	}

	customUserTaskModel struct {
		*defaultUserTaskModel
	}
)

// NewUserTaskModel returns a model for the database table.
func NewUserTaskModel(conn sqlx.SqlConn) UserTaskModel {
	return &customUserTaskModel{
		defaultUserTaskModel: newUserTaskModel(conn),
	}
}
