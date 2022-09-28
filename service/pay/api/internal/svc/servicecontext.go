package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"micro-mall-server/service/pay/api/internal/config"
	"micro-mall-server/service/pay/rpc/pay"
)

type ServiceContext struct {
	Config config.Config

	PayRpc pay.Pay
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: pay.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
