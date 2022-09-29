package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	// http service config.
	rest.RestConf

	// 鉴权
	Auth struct {
		// 密钥
		AccessSecret string
		// 过期时间
		AccessExpire int64
	}

	// rpc 客户端配置
	UserRpc zrpc.RpcClientConf
}
