package svc

import (
	"backend/internal/config"
	"backend/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config      config.Config
	AuthRequest rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AuthRequest: middleware.NewAuthRequestMiddleware().Handle,
	}
}
