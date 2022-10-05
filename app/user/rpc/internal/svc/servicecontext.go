package svc

import (
	"micro-mall-server/app/user/model"
	"micro-mall-server/app/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	UserModel     model.UserModel
	UserAuthModel model.UserAuthModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
