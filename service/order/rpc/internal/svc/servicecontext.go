package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"micro-mall-server/service/order/model"
	"micro-mall-server/service/order/rpc/internal/config"
	"micro-mall-server/service/product/rpc/product"
	"micro-mall-server/service/user/rpc/user"
)

type ServiceContext struct {
	Config config.Config

	OrderModel model.OrderModel

	UserRpc    user.User
	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:     c,
		OrderModel: model.NewOrderModel(conn, c.CacheRedis),
		UserRpc:    user.NewUser(zrpc.MustNewClient(c.UserRpc)),
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
