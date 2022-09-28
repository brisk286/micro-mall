package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"micro-mall-server/service/product/api/internal/config"
	"micro-mall-server/service/product/rpc/product"
)

type ServiceContext struct {
	Config config.Config

	ProductRpc product.Product
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		ProductRpc: product.NewProduct(zrpc.MustNewClient(c.ProductRpc)),
	}
}
