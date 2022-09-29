package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"micro-mall-server/service/user/model"
	"micro-mall-server/service/user/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config

	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// MySQL 连接
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
