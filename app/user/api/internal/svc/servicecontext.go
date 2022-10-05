package svc

import (
	"micro-mall-server/app/user/api/internal/config"
	"micro-mall-server/app/user/rpc/userrpc"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc userrpc.Userrpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
