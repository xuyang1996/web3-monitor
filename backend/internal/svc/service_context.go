package svc

import (
	"backend/internal/config"
	"backend/internal/middleware"
	"backend/internal/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config           config.Config
	AuthRequest      rest.Middleware
	UserModel        model.UserModel
	UserTaskModel    model.UserTaskModel
	TransferTxnModel model.TransferTxnModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)

	return &ServiceContext{
		Config:           c,
		AuthRequest:      middleware.NewAuthRequestMiddleware().Handle,
		UserModel:        model.NewUserModel(sqlConn),
		UserTaskModel:    model.NewUserTaskModel(sqlConn),
		TransferTxnModel: model.NewTransferTxnModel(sqlConn),
	}
}
