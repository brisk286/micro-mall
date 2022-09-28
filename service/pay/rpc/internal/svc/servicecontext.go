package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"micro-mall-server/service/order/rpc/order"
	"micro-mall-server/service/pay/model"
	"micro-mall-server/service/pay/rpc/internal/config"
	"micro-mall-server/service/user/rpc/user"
)

type ServiceContext struct {
	Config config.Config

	PayModel model.PayModel

	UserRpc  user.User
	OrderRpc order.Order
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:   c,
		PayModel: model.NewPayModel(conn, c.CacheRedis),
		UserRpc:  user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: order.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
